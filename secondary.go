package goline

import (
	"bytes"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/shillbie/register/LineThrift"
	//"strconv"
)

func encryptCredential(email, pwd string, key *LineThrift.RSAKey) (string, error) {
	decN, err := hex.DecodeString(key.Nvalue)
	if err != nil {
		return "", err
	}
	n := big.NewInt(0)
	n.SetBytes(decN)
	eVal, err := strconv.ParseInt(key.Evalue, 16, 32)
	if err != nil {
		return "", err
	}
	e := int(eVal)
	rsaKey := rsa.PublicKey{N: n, E: e}
	msg := []byte(string(rune(len(key.SessionKey))) + key.SessionKey + string(rune(len(email))) + email + string(rune(len(pwd))) + pwd)
	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, &rsaKey, msg)
	return hex.EncodeToString(cipherText), err
}

func (p *LINE) SetEmailIdentifier(email, passwd string) error {
	auth := p.VerifyService()
	session, err := auth.OpenAuthSession(p.ctx, &LineThrift.AuthSessionRequest{MetaData: map[string]string{}})
	if err != nil {
		return err
	}
	settings, err := p.GetSettingsAttributes2([]int8{24})
	if err != nil {
		return err
	}
	rsaKey, err := auth.GetAuthRSAKey(p.ctx, session, 1)
	if err != nil {
		return err
	}
	ciphertext, err := encryptCredential(email, passwd, rsaKey)
	if err != nil {
		return err
	}
	if settings.IdentityIdentifier == "" {
		_, err = auth.SetIdentifier(p.ctx, session, &LineThrift.IdentityCredentialRequest{IdentityProvider: 1, CipherKeyId: rsaKey.Keynm, CipherText: ciphertext})
	} else {
		_, err = auth.UpdateIdentifier(p.ctx, session, &LineThrift.IdentityCredentialRequest{IdentityProvider: 1, CipherKeyId: rsaKey.Keynm, CipherText: ciphertext})
	}
	if err != nil {
		return err
	}
	fmt.Print("請輸入驗證碼:")
	var pin = ""
	fmt.Scanln(&pin)
	_, err = auth.ConfirmIdentifier(p.ctx, session, &LineThrift.IdentityCredentialRequest{ConfirmationRequest: &LineThrift.IdentifierConfirmationRequest{ForceRegistration: true, VerificationCode: pin}})
	return err
}

func (p *LINE) SetTempEmailIdentifier() (email string, err error) {

	// timestamp := time.Now().UnixNano()
	// enc := md5.New()
	// enc.Write([]byte(fmt.Sprintf("%d", timestamp)))
	// email := "cawaii+" + hex.EncodeToString(enc.Sum(nil)) + "@logicstreak.com"
	email, err = CreateTempEmail("")
	if err != nil {
		return email, err
	}

	auth := p.VerifyService()
	session, err := auth.OpenAuthSession(p.ctx, &LineThrift.AuthSessionRequest{MetaData: map[string]string{}})
	if err != nil {
		return email, err
	}
	settings, err := p.GetSettingsAttributes2([]int8{24})
	if err != nil {
		return email, err
	}
	rsaKey, err := auth.GetAuthRSAKey(p.ctx, session, 1)
	if err != nil {
		return email, err
	}
	ciphertext, err := encryptCredential(email, line_password, rsaKey)
	if err != nil {
		return email, err
	}
	if settings.IdentityIdentifier == "" {
		_, err = auth.SetIdentifier(p.ctx, session, &LineThrift.IdentityCredentialRequest{IdentityProvider: 1, CipherKeyId: rsaKey.Keynm, CipherText: ciphertext})
	} else {
		_, err = auth.UpdateIdentifier(p.ctx, session, &LineThrift.IdentityCredentialRequest{IdentityProvider: 1, CipherKeyId: rsaKey.Keynm, CipherText: ciphertext})
	}
	if err != nil {
		return email, err
	}
	var pin string
	jwt, err := GetTempmailToken(email)
	if err != nil {
		return email, err
	}
	for i := 0; i < 60; i++ {
		time.Sleep(1 * time.Second)
		res, err := GetTempmailMessage(jwt)
		if err != nil {
			return email, err
		}
		if len(res.Messages) > 0 {
			DeleteTempmailMessage(jwt, res.Messages[0].ID)
			match, _ := regexp.Compile("[0-9][0-9][0-9][0-9]")
			pin = match.FindString(res.Messages[0].Intro)
			break
		}
	}
	_, err = auth.ConfirmIdentifier(p.ctx, session, &LineThrift.IdentityCredentialRequest{ConfirmationRequest: &LineThrift.IdentifierConfirmationRequest{ForceRegistration: true, VerificationCode: pin}})
	return email, err
}

func (p *LINE) UpdatePassword(newPass string) error {
	auth := p.VerifyService()
	session, err := auth.OpenAuthSession(p.ctx, &LineThrift.AuthSessionRequest{MetaData: map[string]string{}})
	if err != nil {
		return err
	}
	settings, err := p.GetSettingsAttributes2([]int8{24})
	if err != nil {
		return err
	}
	rsaKey, err := auth.GetAuthRSAKey(p.ctx, session, 1)
	if err != nil {
		return err
	}
	ciphertext, err := encryptCredential(settings.IdentityIdentifier, newPass, rsaKey)
	res, err := auth.UpdatePassword(p.ctx, session, &LineThrift.IdentityCredentialRequest{IdentityProvider: 1, CipherKeyId: rsaKey.Keynm, CipherText: ciphertext})
	fmt.Println(res, err)
	return err
}

func GetTempDomains() (domain string, err error) {
	req, _ := http.NewRequest("GET", "https://api.mail.tm/domains?page=1", nil)
	req.Header.Set("accept", "application/ld+json")
	client := &http.Client{}
	resp, err := client.Do(req)
	data, _ := ioutil.ReadAll(resp.Body)
	result := struct {
		Member []struct {
			Domain string `json:"domain"`
		} `json:"hydra:member"`
	}{}
	json.Unmarshal(data, &result)
	return result.Member[0].Domain, err
}

func CreateTempEmail(account string) (mail string, err error) {
	domain, err := GetTempDomains()
	if err != nil {
		return "", err
	}
	if account == "" {
		timestamp := time.Now().UnixNano()
		enc := md5.New()
		enc.Write([]byte(fmt.Sprintf("%d", timestamp)))
		account = hex.EncodeToString(enc.Sum(nil))
	}
	client := &http.Client{}
	body := &bytes.Buffer{}
	content, _ := json.Marshal(map[string]string{
		"address":  account + "@" + domain,
		"password": account,
	})
	body.Write(content)
	req, _ := http.NewRequest("POST", "https://api.mail.tm/accounts", body)
	req.Header.Set("accept", "application/ld+json")
	req.Header.Set("Content-Type", "application/json")
	_, err = client.Do(req)
	return account + "@" + domain, err
}

func GetTempmailToken(account string) (jwtToken string, err error) {
	client := &http.Client{}
	body := &bytes.Buffer{}
	content, _ := json.Marshal(map[string]string{
		"address":  account,
		"password": strings.Split(account, "@")[0],
	})
	body.Write(content)
	req, _ := http.NewRequest("POST", "https://api.mail.tm/token", body)
	req.Header.Set("accept", "application/ld+json")
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	data, _ := ioutil.ReadAll(resp.Body)
	result := map[string]string{"token": ""}
	json.Unmarshal(data, &result)
	return result["token"], err
}

type tempMailResult struct {
	Messages []struct {
		ID        string `json:"id"`
		Subject   string `json:"subject"`
		Intro     string `json:"intro"`
		CreatedAt string `json:"createdAt"`
		UpdatedAt string `json:"updatedAt"`
	} `json:"hydra:member"`
}

func GetTempmailMessage(jwtToken string) (result tempMailResult, err error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://api.mail.tm/messages?page=1", nil)
	req.Header.Set("accept", "application/ld+json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+jwtToken)
	resp, err := client.Do(req)
	if err != nil {
		return result, err
	}
	data, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(data, &result)
	return result, err
}

func DeleteTempmailMessage(jwtToken, id string) error {
	client := &http.Client{}
	req, _ := http.NewRequest("DELETE", "https://api.mail.tm/messages/"+id, nil)
	req.Header.Set("accept", "application/ld+json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+jwtToken)
	_, err := client.Do(req)
	return err
}
