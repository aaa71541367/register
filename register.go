package goline

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/shillbie/register/LineThrift"
	"github.com/shillbie/register/helper"
	"github.com/shillbie/register/thrift"

	"github.com/ericlagergren/siv"
	"github.com/google/uuid"
)

var debug bool = false
var captchaSolver = ""

func (p *LINE) RegisterService(country string) *LineThrift.PrimaryAccountInitServiceClient {
	ip, err := helper.GetRandomIP(country)
	if err != nil {
		p.Proxy = helper.ZoneRegister.GetPorxyByCountry(country)
	} else {
		p.Proxy = helper.ZoneProxy1.GetPorxyByCountry(country)
	}
	p.setDefaultHttpClient()
	httpClient, _ := thrift.NewTHttpClientWithOptions(LINE_HOST_DOMAIN+"/acct/pais/v1", thrift.THttpClientOptions{
		Client: p.connection,
	})
	// httpClient, _ := thrift.NewTHttpClient("https://gw.line.naver.jp/acct/pais/v1")
	p.Proxy = "fake://" + ip.String() + ":8080"
	buffer := thrift.NewTBufferedTransportFactory(4096)
	trans := httpClient.(*thrift.THttpClient)
	trans.SetHeader("User-Agent", p.userAgent)
	trans.SetHeader("X-Line-Application", p.AppName)
	trans.SetHeader("Connection", "Keep-Alive")
	trans.SetHeader("x-lal", "zh-Hant_TW")
	trans.SetHeader("x-lpv", "1")
	if err == nil {
		trans.SetHeader("x-forwarded-for", ip.String())
	}
	buftrans, _ := buffer.GetTransport(trans)
	compactProtocol := thrift.NewTCompactProtocolFactory()
	return LineThrift.NewPrimaryAccountInitServiceClientFactory(buftrans, compactProtocol)
}

func generateUUID(device int) string {
	if device == 0 {
		random := uuid.New()
		return strings.ReplaceAll(random.String(), "-", "")
	} else {
		random := uuid.New()
		return strings.ReplaceAll(random.String(), "-", "")[:25]
	}
}

func (p *LINE) RegisterByPhone(country string, sms helper.SmsReciever) (token string, err error) {
	device := &LineThrift.Device{Udid: generateUUID(0), DeviceModel: "SM-" + []string{"A125F", "X200", "T295", "A5360", "A528B", "N960F", "N950F", "N770F", "T510", "A7050"}[rand.Intn(10)]}
	appVer := []string{"13.1.0", "13.1.2", "13.1.3"}[rand.Intn(3)]
	deviceVer := []string{"12.6.1", "9.0.0", "8.1.0", "10.0.0"}[rand.Intn(4)]
	p.userAgent = "Line/" + appVer
	p.AppName = "ANDROID\t" + appVer + "\tAndroid OS\t" + deviceVer + ";MA_UNV"
	auth := p.AuthService()
	auth.GetAnalyticsInfo(p.ctx)
	auth.NotifyInstalled(p.ctx, device.Udid, p.AppName)
	phoneNum, orderid, err := sms.GetPhoneNum(strings.ToLower(country))
	if err != nil {
		return "GetPhoneNum", err
	}
	fmt.Println("New Phone Number:", phoneNum)
	var register *LineThrift.PrimaryAccountInitServiceClient
	register = p.RegisterService(strings.ToLower(country))
	session, err := register.OpenSession(context.TODO(), &LineThrift.OpenSessionRequest{})
	if err != nil {
		return "OpenSession", err
	}
	fmt.Println("Your Auth Session:", session)
	phone := &LineThrift.UserPhoneNumber{PhoneNumber: phoneNum, CountryCode: strings.ToUpper(country)}
	register.GetCountryInfo(context.TODO(), session, &LineThrift.SimCard{CountryCode: strings.ToUpper(country), Hni: helper.GetSimHni(country), CarrierName: helper.GetSimCarrier(country)})
	// register.LookupAvailableEap(context.TODO(), &LineThrift.LookupAvailableEapRequest{AuthSessionId: session})
RequestPin:
	methods, err := register.GetPhoneVerifMethodV2(context.TODO(), &LineThrift.GetPhoneVerifMethodV2Request{AuthSessionId: session, Device: device, UserPhoneNumber: phone})
	if err != nil {
		return "GetPhoneVerifMethod", err
	}
	if _, err := register.RequestToSendPhonePinCode(context.TODO(), &LineThrift.ReqToSendPhonePinCodeRequest{AuthSessionId: session, UserPhoneNumber: phone, VerifMethod: LineThrift.VerifMethod(methods.AvailableMethods[0])}); err != nil {
		if authex, ok := err.(*LineThrift.AuthException); ok && authex.Code == LineThrift.AuthErrorCode_HUMAN_VERIFICATION_REQUIRED {
			fmt.Println("HUMAN_VERIFICATION_REQUIRED: ", authex.WebAuthDetails)
			token, err := solve(authex.WebAuthDetails)
			if err != nil {
				return "solveCaptcha", err
			}
			fmt.Println("HUMAN_VERIFICATION_SOLVED: ", token)
			err = p.sendGoogleRecaptchaTokenToLine(authex.WebAuthDetails, token)
			if err != nil {
				return "sendGoogleRecaptchaTokenToLine", err
			}
			goto RequestPin
		} else {
			sms.ReleasePhone(orderid, "shieldNumber")
			return "SendPinCodeForPhone", err
		}
	}
	if _, err := register.VerifyPhonePinCode(context.TODO(), &LineThrift.VerifyPhonePinCodeRequest{AuthSessionId: session, UserPhoneNumber: phone, PinCode: "000000"}); err != nil {
		if authex, ok := err.(*LineThrift.AuthException); ok && authex.Code == LineThrift.AuthErrorCode_INVALID_CONTEXT {
			fmt.Println(authex)
			sms.ReleasePhone(orderid, "shieldNumber")
			return "TestVerifyPinCode", fmt.Errorf("error verify pincode")
		}
	}
	// fmt.Println(register.FetchPhonePinCodeMsg(context.TODO(), &LineThrift.FetchPhonePinCodeMsgRequest{AuthSessionId: session, UserPhoneNumber: phone}))
	code, err := sms.GetPhoneCode(orderid, 30)
	if err != nil {
		if _, err := register.RequestToSendPhonePinCode(context.TODO(), &LineThrift.ReqToSendPhonePinCodeRequest{AuthSessionId: session, UserPhoneNumber: phone, VerifMethod: LineThrift.VerifMethod(methods.AvailableMethods[0])}); err != nil {
			sms.ReleasePhone(orderid, "releaseNumber")
			return "SendPinCodeForPhone", err
		}
		code, err = sms.GetPhoneCode(orderid, 10)
	}
	if code == "" || err != nil {
		sms.ReleasePhone(orderid, "shieldNumber")
		return "GetPinCodeFromPhone", fmt.Errorf("can't recieve pincode")
	}
	fmt.Println("Got Pin Code:", code)
VerifyPin:
	if _, err := register.VerifyPhonePinCode(context.TODO(), &LineThrift.VerifyPhonePinCodeRequest{AuthSessionId: session, UserPhoneNumber: phone, PinCode: code}); err != nil {
		if authex, ok := err.(*LineThrift.AuthException); ok && authex.Code == LineThrift.AuthErrorCode_HUMAN_VERIFICATION_REQUIRED {
			fmt.Println("HUMAN_VERIFICATION_REQUIRED: ", authex.WebAuthDetails)
			token, err := solve(authex.WebAuthDetails)
			if err != nil {
				return "solveCaptcha", err
			}
			fmt.Println("HUMAN_VERIFICATION_SOLVED: ", token)
			err = p.sendGoogleRecaptchaTokenToLine(authex.WebAuthDetails, token)
			if err != nil {
				return "sendGoogleRecaptchaTokenToLine", err
			}
			goto VerifyPin
		} else {
			sms.ReleasePhone(orderid, "shieldNumber")
			return "VerifyPhonePin", err
		}
	}
	name := generateUUID(0)[:7]
	if err := register.ValidateProfile(context.TODO(), session, name); err != nil {
		return "ValidateProfile", err
	}
SendRegister:
	keypair := GenerateAsymmetricKeypair()
	nonce := randomBytes(16)
	exchangedData, err := register.ExchangeEncryptionKey(context.TODO(), session, &LineThrift.ExchangeEncryptionKeyRequest{AuthKeyVersion: 1, PublicKey: string(base64.StdEncoding.EncodeToString(keypair.PublicKey)), Nonce: string(base64.StdEncoding.EncodeToString(nonce))})
	if err != nil {
		return "ExchangeEncryptionKey", err
	}
	cipherText, err := EncryptPassword(line_password, keypair, nonce, exchangedData.PublicKey, exchangedData.Nonce)
	if err != nil {
		return "EncryptPassword", err
	}
	if err := register.SetPassword(context.TODO(), session, &LineThrift.EncryptedPassword{EncryptionKeyVersion: 1, CipherText: cipherText}); err != nil {
		return "SetPassword", err
	}
	res, err := register.RegisterPrimaryUsingPhoneWithTokenV3(context.TODO(), session)
	if err != nil {
		if authex, ok := err.(*LineThrift.AuthException); ok && authex.Code == LineThrift.AuthErrorCode_HUMAN_VERIFICATION_REQUIRED {
			fmt.Println("HUMAN_VERIFICATION_REQUIRED: ", authex.WebAuthDetails)
			token, err := solve(authex.WebAuthDetails)
			if err != nil {
				return "solveCaptcha", err
			}
			fmt.Println("HUMAN_VERIFICATION_SOLVED: ", token)
			err = p.sendGoogleRecaptchaTokenToLine(authex.WebAuthDetails, token)
			if err != nil {
				return "sendGoogleRecaptchaTokenToLine", err
			}
			goto SendRegister
		}
	}
	sms.ReleasePhone(orderid, "releaseNumber")
	fmt.Println("• Date Making: ", time.Now().Format("2006-01-02"))
	fmt.Println("• AuthKey: ", res.AuthToken)
	fmt.Println("• Token V3: ", res.TokenV3IssueResult_.AccessToken)
	fmt.Println("• RefreshToken: ", res.TokenV3IssueResult_.RefreshToken)
	fmt.Println("• Pw: ", line_password)
	primary := GetSignature(res.AuthToken, 0)
	fmt.Println("• Primary: ", primary)
	p.connection = nil

	p.AuthToken = res.TokenV3IssueResult_.AccessToken
	p.UpdateProfileAttributes(map[int32]*LineThrift.ProfileContent{2: {Value: name}})
	if debug {
		fmt.Println("按Enter鍵繼續...")
		os.Stdin.Read(make([]byte, 10))
	}
	p.TalkService().GetPendingAgreements(p.ctx)
	if debug {
		fmt.Println("按Enter鍵繼續...")
		os.Stdin.Read(make([]byte, 10))
	}
	prof, _ := p.GetProfile()
	p.MID = prof.Mid
	if debug {
		fmt.Println("按Enter鍵繼續...")
		os.Stdin.Read(make([]byte, 10))
	}

	p.GetSettings()
	if debug {
		fmt.Println("按Enter鍵繼續...")
		os.Stdin.Read(make([]byte, 10))
	}
	p.UpdateSettingsAttributes2(&LineThrift.Settings{PrivacySearchByPhoneNumber: true, PrivacySyncContacts: true}, []LineThrift.SettingsAttribute{6, 7})
	if debug {
		fmt.Println("按Enter鍵繼續...")
		os.Stdin.Read(make([]byte, 10))
	}
	p.GetConfigurations(strings.ToUpper(country))
	if debug {
		fmt.Println("按Enter鍵繼續...")
		os.Stdin.Read(make([]byte, 10))
	}
	key := GenerateAsymmetricKeypair()
	keyData, _ := p.TalkService().RegisterE2EEPublicKey(context.TODO(), 0, &LineThrift.E2EEPublicKey{Version: 1, KeyId: 0, KeyData: key.PublicKey, CreatedTime: time.Now().Unix() * 1000})
	p.E2EEKey = &LineThrift.E2EEKey{
		Version:     keyData.Version,
		KeyId:       keyData.KeyId,
		PublicKey:   keyData.KeyData,
		PrivateKey:  key.PrivateKey,
		CreatedTime: keyData.CreatedTime,
	}
	fmt.Println("• E2EE_KEY ID: ", keyData.KeyId)
	fmt.Println("• E2EE_KEY Public: ", base64.StdEncoding.EncodeToString(keyData.KeyData))
	fmt.Println("• E2EE_KEY Private: ", base64.StdEncoding.EncodeToString(key.PrivateKey))
	if debug {
		fmt.Println("按Enter鍵繼續...")
		os.Stdin.Read(make([]byte, 10))
	}
	p.GetAllContactIds()
	if debug {
		fmt.Println("按Enter鍵繼續...")
		os.Stdin.Read(make([]byte, 10))
	}
	// p.GetAllChatMids(true, true)
	// if debug {
	// 	fmt.Println("按Enter鍵繼續...")
	// 	os.Stdin.Read(make([]byte, 10))
	// }
	// p.AcquireEncryptedAccessToken()
	// p.pushClient = p.PushService()
	// if debug {
	// 	fmt.Println("按Enter鍵繼續...")
	// 	os.Stdin.Read(make([]byte, 10))
	// }
	// p.Sync()
	// p.NotifyRegistrationComplete(device.Udid)
	// if debug {
	// 	fmt.Println("按Enter鍵繼續...")
	// 	os.Stdin.Read(make([]byte, 10))
	// }
	return primary, nil
}

func (p *LINE) RegisterE2EEKey() error {
	if p.E2EEKey != nil {
		return nil
	} else if strings.HasPrefix(p.AuthToken, "u") {
		myKeys, err := p.TalkService().GetE2EEPublicKeys(context.TODO())
		if err != nil {
			return err
		}
		for _, key := range myKeys {
			err = p.TalkService().RemoveE2EEPublicKey(context.TODO(), key)
			if err != nil {
				return err
			}
		}
		key := GenerateAsymmetricKeypair()
		res, err := p.TalkService().RegisterE2EEPublicKey(context.TODO(), 0, &LineThrift.E2EEPublicKey{Version: 1, KeyId: 0, KeyData: key.PublicKey, CreatedTime: time.Now().Unix() * 1000})
		if err != nil {
			return err
		}
		p.E2EEKey = &LineThrift.E2EEKey{
			Version:     res.Version,
			KeyId:       res.KeyId,
			PublicKey:   res.KeyData,
			PrivateKey:  key.PrivateKey,
			CreatedTime: res.CreatedTime,
		}

		bData, _ := json.Marshal(p.E2EEKey)
		_, err = p.CreateKeep("account_e2ee_key", "e2ee_key: "+base64.StdEncoding.EncodeToString(bData))
		return err
	}
	return fmt.Errorf("reregister key error.")
}

func (p *LINE) sendGoogleRecaptchaTokenToLine(details *LineThrift.WebAuthDetails, token string) error {
	jsonStr, err := json.Marshal(map[string]string{"verifier": token})
	if err != nil {
		return err
	}
	request, err := http.NewRequest("POST", LINE_W_DOMAIN+"/sec/v3/recaptcha/result/verify", bytes.NewBuffer(jsonStr))
	if err != nil {
		return err
	}
	for k, v := range map[string]string{
		"Connection":       "keep-alive",
		"user-agent":       "Mozilla/5.0 (Linux; Android 8.0.0; Nexus 6P Build/OPR6.170623.019; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/88.0.4324.93 Mobile Safari/537.36",
		"Accept":           "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
		"Referer":          LINE_W_DOMAIN + "/sec/v3/recaptcha",
		"Origin":           LINE_W_DOMAIN,
		"X-Requested-With": "jp.naver.line.android",
		"Content-Type":     "application/json;charset=UTF-8",
		"Host":             "w.line.me",
	} {
		request.Header.Set(k, v)
	}
	request.AddCookie(&http.Cookie{
		Name:  "lsct_acct_init",
		Value: details.Token[15:],
	})
	response, err := p.connection.Do(request)
	if err != nil {
		return err
	}
	if response.StatusCode == 200 {
		return nil
	}
	return fmt.Errorf("failed to solve human verify process")

}

func solve(details *LineThrift.WebAuthDetails) (string, error) {
	headers := map[string]string{
		"Connection":                "keep-alive",
		"Upgrade-Insecure-Requests": "1",
		"user-agent":                "Mozilla/5.0 (Linux; Android 8.0.0; Nexus 6P Build/OPR6.170623.019; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/88.0.4324.93 Mobile Safari/537.36",
		"Accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
		"authorization":             details.Token,
		// "Accept-Encoding":           "gzip, deflate",
		"Accept-Language": "ja-JP,ja;q=0.9,en-US;q=0.8,en;q=0.7",
		"Host":            "2captcha.com",
	}
	params := map[string]string{
		"method":      "userrecaptcha",
		"pageurl":     details.BaseUrl,
		"soft_id":     "3029",
		"googlekey":   "6Lfo_XYUAAAAAFQbdsuk6tETqnpKIg5gNxJy4xM0",
		"key":         "37fd7d9df1cbbadd8206f5e00b500b51",
		"header_acao": "1",
	}
	request, err := http.NewRequest("GET", "https://2captcha.com/in.php", nil)
	if err != nil {
		return "", err
	}
	for k, v := range headers {
		request.Header.Set(k, v)
	}
	rParam := request.URL.Query()
	for k, v := range params {
		rParam.Add(k, v)
	}
	request.URL.RawQuery = rParam.Encode()
	response, err := (&http.Client{}).Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	if string(body)[:2] != "OK" {
		return "", fmt.Errorf("something went wrong on starting solve capture process" + string(body))
	}
	captchaId := string(body)[3:]
	for i := 0; i < 40; i++ {
		time.Sleep(time.Second * 5)
		request, err = http.NewRequest("GET", "https://2captcha.com/res.php?key="+captchaSolver+"&action=get&id="+captchaId, nil)
		if err != nil {
			return "", err
		}
		response, err = (&http.Client{}).Do(request)
		if err != nil {
			if strings.Contains(err.Error(), "timeout awaiting") {
				continue
			}
			return "", err
		}
		body, err = ioutil.ReadAll(response.Body)
		if err != nil {
			return "", err
		}
		if string(body)[:2] == "OK" {
			return string(body)[3:], nil
		}
		if string(body) == "CAPCHA_NOT_READY" {
			continue
		}
		response.Body.Close()
	}
	return "", fmt.Errorf("failed solve recaptcha")
}

func (p *LINE) StartMigration() error {
	service := p.PrimaryQRMigrationPreService()
	session, err := service.CreateQRMigrationSession(p.ctx, &LineThrift.CreateSessionRequest{})
	if err != nil {
		return err
	}
	nonce := randomBytes(32)
	keypair := GenerateAsymmetricKeypair()
	qrmap := map[string]string{
		"si": session.SessionId,
		"qi": fmt.Sprintf("%X", nonce),
		"pk": fmt.Sprintf("%X", keypair.PublicKey),
	}
	fmt.Println(session.SessionId, fmt.Sprintf("%X", nonce), fmt.Sprintf("%X", keypair.PublicKey), fmt.Sprintf("%X", keypair.PrivateKey))
	bData, err := json.Marshal(qrmap)
	fmt.Println("https://chart.googleapis.com/chart?chs=500x500&cht=qr&chl=" + url.QueryEscape(string(bData)))
	// qrterminal.Generate(string(bData), qrterminal.L, os.Stdout)
	// go p.TryMigration(qrmap)
	fmt.Println("按Enter鍵繼續...")
	os.Stdin.Read(make([]byte, 10))
	time.Sleep(2 * time.Second)
	service.SendEncryptedE2EEKey(p.ctx, &LineThrift.SendEncryptedE2EEKeyRequest{
		SessionId: session.SessionId,
		EncryptedSecureChannelPayload: &LineThrift.EncryptedBinary{
			RecoveryKey:       []byte{146, 1, 196, 16, 9, 231, 32, 118, 225, 45, 12, 22, 26, 217, 192, 116, 198, 36, 121, 74},
			BackupBlobPayload: []byte{147, 1, 146, 146, 1, 206, 0, 60, 0, 37, 145, 2, 196, 239, 81, 173, 78, 182, 210, 207, 157, 196, 91, 218, 171, 171, 70, 64, 153, 100, 235, 1, 120, 95, 212, 153, 64, 62, 56, 140, 114, 38, 57, 36, 64, 227, 12, 10, 55, 48, 76, 185, 211, 36, 231, 144, 78, 230, 35, 59, 180, 84, 170, 52, 3, 34, 94, 89, 189, 13, 240, 207, 128, 37, 116, 140, 190, 198, 162, 82, 77, 44, 72, 4, 56, 6, 145, 205, 8, 57, 180, 98, 88, 85, 17, 247, 34, 128, 166, 115, 11, 101, 34, 49, 146, 176, 250, 48, 12, 220, 229, 232, 184, 113, 209, 201, 38, 204, 209, 113, 218, 159, 208, 61, 16, 191, 63, 125, 122, 181, 31, 186, 55, 155, 213, 111, 78, 77, 46, 80, 95, 199, 104, 101, 169, 204, 99, 237, 52, 223, 75, 193, 159, 86, 238, 117, 132, 244, 204, 102, 241, 90, 118, 124, 153, 161, 160, 215, 178, 50, 221, 16, 136, 187, 165, 17, 132, 247, 116, 24, 157, 253, 21, 133, 115, 127, 84, 242, 175, 17, 63, 145, 224, 73, 235, 200, 18, 44, 115, 191, 108, 107, 247, 89, 82, 238, 106, 87, 253, 181, 101, 171, 175, 33, 176, 32, 227, 92, 190, 107, 99, 108, 160, 251, 50, 167, 232, 160, 26, 211, 221, 54, 156, 51, 217, 125, 223, 23, 187, 210, 154, 147, 22, 225, 17, 119, 39, 103, 15, 98, 151, 58, 92},
		},
	})
	return nil
}

func (p *LINE) TryMigration(data map[string]string) (string, error) {
	nb := NewLogin()
	session := data["si"]
	identifier, _ := hex.DecodeString(data["qi"])
	nonce := randomBytes(12)
	pubkey, _ := hex.DecodeString(data["pk"])
	nb.AuthToken = session
	keypair := GenerateAsymmetricKeypair()
	sharedSecret, err := GenerateSharedKey(keypair.PrivateKey, pubkey)
	if err != nil {
		return "", err
	}
	aead, err := siv.NewGCM(sharedSecret)
	if err != nil {
		return "", err
	}
	ciphertext := aead.Seal(nil, nonce, identifier, nil)
	vservice := nb.PrimaryQRMigrationLPService()
	resp, err := vservice.CheckIfEncryptedE2EEKeyReceived(p.ctx, &LineThrift.CheckIfEncryptedE2EEKeyReceivedRequest{
		SessionId: session,
		SecureChannelData: &LineThrift.SecureChannelData{
			NewDevicePublicKey_:   keypair.PublicKey,
			EncryptedQrIdentifier: fmt.Sprintf("%X", append(nonce, ciphertext...)),
		},
	})
	if err != nil {
		return "", err
	}
	service := nb.PrimaryQRMigrationService()
	res, err := service.MigratePrimaryUsingQrCode(nb.ctx, &LineThrift.MigratePrimaryUsingQrCodeRequest{
		SessionId: session,
		Nonce:     resp.Nonce,
		NewDevice_: &LineThrift.Device2{
			DeviceModel: "iPhone13,2",
			Udid:        "unknown" + generateUUID(1),
		},
	})
	fmt.Println(res)
	if err == nil {
		return GetSignature(res.Mid+":"+res.TokenV1IssueResult_.TokenSecret, 0), nil
	}
	return "", err
}
