package goline

import (
	"fmt"

	"github.com/shillbie/register/LineThrift"

	//"time"
	"encoding/base64"
	"encoding/json"

	"github.com/tidwall/gjson"
	//    "crypto/sha256"
	//    "encoding/base64"
)

type Results struct {
	Result struct {
		Verifier  string `json:"verifier"`
		AuthPhase string `json:"authPhase"`
		Metadata  struct {
			EncryptedKeyChain string `json:"encryptedKeyChain"`
			PublicKey         string `json:"publicKey"`
			HashKeyChain      string `json:"hashKeyChain"`
		} `json:"matadata"`
	} `json:"result"`
	Timestamp string `json:"timestamp"`
}

func getResult(body []byte) (*Results, error) {
	var s = new(Results)
	err := json.Unmarshal(body, &s)
	return s, err
}

func (p *LINE) LoadService(showLog bool) (err error) {
	getprof, err := p.GetProfile()
	if err != nil {
		return err
	}
	//contact, err := p.GetContact(p.MID)
	// if err != nil {
	// 	return err
	// }
	p.MID = getprof.Mid
	if showLog {
		fmt.Println("\n##########  -[Login Successful]-  ##########")
		fmt.Println("DisplayName: ", getprof.DisplayName)
		fmt.Println("MID: ", p.MID)
		fmt.Println("AuthToken: ", p.AuthToken)
		fmt.Println("########## -[End of Information]- ##########")
		// p.SetTimelineHeaders()
		//s.AcquireEncryptedAccessToken()
	}
	p.Friends, err = p.GetAllContactIds()
	if err != nil {
		return err
	}
	p.Revision, err = p.PollService().GetLastOpRevision(p.ctx)
	// p.Revision = -1
	// p.Revision = 0
	if err != nil {
		return err
	}
	err = p.LoadE2EEKeys()
	return err
}

func (p *LINE) LoadE2EEKeys() error {
	if keeps, err := p.GetKeeps("0", "30"); err == nil {
		if gjson.Get(keeps, "message").String() == "success" {
			pubKeys, err := p.GetE2EEPublicKeys()
			if err != nil || len(pubKeys) < 1 {
				return err
			}
			for _, keep := range gjson.Get(keeps, `result.contents.#(contentData.#(text%"e2ee_key: *")).contentData`).Array() {
				bData, _ := base64.StdEncoding.DecodeString(keep.Get("text").String()[10:])
				var myKey *LineThrift.E2EEKey
				json.Unmarshal(bData, &myKey)
				if myKey.KeyId == pubKeys[0].KeyId {
					p.E2EEKey = myKey
					fmt.Println(p.E2EEKey)
				}
			}
			return nil
		}
	} else {
		return err
	}
	return fmt.Errorf("unknown error")
}

func (p *LINE) LoginWithAuthToken(authToken string) error {
	p.AuthToken = authToken
	return p.LoadService(true)
}

/*

func loginRequestQR(identity LineThrift.IdentityProvider, verifier string, secret []byte, e2ee int32) *LineThrift.LoginRequest {
	lreq := &LineThrift.LoginRequest{
		Type:             1,
		KeepLoggedIn:     true,
		IdentityProvider: identity,
		AccessLocation:   "127.0.0.1",
		SystemName:       SYSTEM_NAME,
		Verifier:         verifier,
		Secret:           secret,
		E2eeVersion:      e2ee,
	}
	return lreq
}

func (p *LINE) LoginWithQrCode(writeToFile bool) {
	tauth := p.AuthService()
	qrCode, err := tauth.GetAuthQrcode(p.ctx, true, SYSTEM_NAME, true)
	if err != nil {
		panic(err)
	}

	// by jay
	if writeToFile {
		fo, err := os.Create("url_login.txt")
		if err == nil {
			ss := qrCode.Verifier
			buf := make([]byte, 1024)
			buf = []byte(ss)
			_, err := fo.Write(buf[0:len(ss)])
			if err == nil {
				fo.Close()
			}
		}
	}
	p_key := GenerateAsymmetricKeypair()
	secret_query := CreateSecretQuery(p_key.PublicKey)
	fmt.Println("line://au/q/" + qrCode.Verifier + "?secret=" + secret_query + "&e2eeVersion=1")
	client := &http.Client{}
	req, _ := http.NewRequest("GET", LINE_HOST_DOMAIN+LINE_CERTIFICATE_PATH, nil)
	req.Header.Set("User-Agent", p.userAgent)
	req.Header.Set("X-Line-Application", p.AppName)
	req.Header.Set("X-Line-Access", qrCode.Verifier)
	p.AuthToken = qrCode.Verifier
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	x, _ := getResult([]byte(body))
	iR := x.Result
	_verifier := iR.Verifier
	loginZ := p.LoginZService()
	loginReq := loginRequestQR(1, _verifier, []byte{}, 0)
	resultz, err := loginZ.LoginZ(p.ctx, loginReq)
	if err != nil {
		panic(err)
	}
	p.LoginWithAuthToken(resultz.AuthToken)
}
*/
