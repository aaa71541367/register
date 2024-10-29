package goline

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/shillbie/register/LineThrift"
	"github.com/shillbie/register/helper"
	//"fmt"
	//"encoding/json"
)

func (p *LINE) AllowLiff() error {
	url := LINE_ACCESS_DOMAIN + "/dialog/api/permissions"
	data := map[string][]string{"on": {"P", "CM"}, "off": {}}
	byteData, _ := json.Marshal(data)
	extraHeader := map[string]string{
		"X-Line-ChannelId": LINE_LIFF_ID[:10],
		"Content-Type":     "application/json",
	}
	_, err := p.PostContent(url, byteData, extraHeader)
	return err
}

func (p *LINE) AllowLiffConsent(consentUrl string) error {
	client := &http.Client{}
	if p.Proxy != "" {
		proxyURL, _ := url.Parse(p.Proxy)
		client = &http.Client{Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		}}
	}
	req, _ := http.NewRequest("GET", consentUrl, nil)
	req.Header.Set("X-Line-Access", p.AuthToken)
	req.Header.Set("X-Line-Application", p.AppName)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Linux; Android 8.0.1; SAMSUNG Realise/DeachSword; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/56.0.2924.87 Mobile Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	bData, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	channelIds, ok := helper.FormatCommand(string(bData), "<input type=\"hidden\" name=\"channelId\" value=\"[TXT]\">")
	csrfs, ok2 := helper.FormatCommand(string(bData), "<input type=\"hidden\" name=\"__csrf\" id=\"__csrf\" value=\"[TXT]\">")
	if ok && ok2 && len(channelIds) == 1 && len(csrfs) == 1 {
		consentUrl = LINE_ACCESS_DOMAIN + "/oauth2/v2.1/authorize/consent"
		params := url.Values{}
		params.Add("allPermission", "P")
		params.Add("allPermission", "CM")
		params.Add("approvedPermission", "P")
		params.Add("approvedPermission", "CM")
		params.Add("channelId", channelIds[0])
		params.Add("__csrf", csrfs[0])
		params.Add("__WLS", "")
		params.Add("allow", "True")
		req, _ = http.NewRequest("POST", consentUrl, bytes.NewBuffer([]byte(params.Encode())))
		req.Header.Set("X-Line-Access", p.AuthToken)
		req.Header.Set("X-Line-Application", p.AppName)
		req.Header.Set("User-Agent", "Mozilla/5.0 (Linux; Android 8.0.1; SAMSUNG Realise/DeachSword; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/56.0.2924.87 Mobile Safari/537.36")
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		for _, cookie := range resp.Cookies() {
			req.AddCookie(cookie)
		}
		resp, err = client.Do(req)
		if resp.StatusCode != 200 {
			return fmt.Errorf("Can't allow liff")
		}
		return nil
	}
	return fmt.Errorf("Can't find channelID or csrf token")
}

func (p *LINE) SendFlex(to string, data string) error {
	accessToken := ""
	if token, ok := p.liffTokens[to]; !ok {
		token, err := p.LiffService().IssueLiffView(p.ctx, &LineThrift.LiffViewRequest{LiffId: LINE_LIFF_ID, Context: &LineThrift.LiffContext{Chat: &LineThrift.LiffChatContext{ChatMid: to}}})
		if liffexcption, ok := err.(*LineThrift.LiffException); ok && liffexcption.Code == 3 {
			err = p.AllowLiff()
			if err != nil {
				return err
			}
			token, err = p.LiffService().IssueLiffView(p.ctx, &LineThrift.LiffViewRequest{LiffId: LINE_LIFF_ID, Context: &LineThrift.LiffContext{Chat: &LineThrift.LiffChatContext{ChatMid: to}}})
			if err != nil {
				return err
			}
			fmt.Println(token.AccessToken)
		}
		accessToken = token.AccessToken
	} else {
		accessToken = token
	}
	client := &http.Client{}
	if p.Proxy != "" {
		proxyURL, _ := url.Parse(p.Proxy)
		client = &http.Client{Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		}}
	}
	url := LINE_API_DOMAIN + "/message/v3/share"
	datas := `{"messages": [` + data + `]}`
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte(datas)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+accessToken)
	res, err := client.Do(req)
	if err == nil {
		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)
		fmt.Println(string(body))
	}
	return err
}

func (p *LINE) SendSquareFlex(to string, data string) error {
	accessToken := ""
	if token, ok := p.liffTokens[to]; !ok {
		token, err := p.LiffService().IssueLiffView(p.ctx, &LineThrift.LiffViewRequest{LiffId: LINE_LIFF_ID, Context: &LineThrift.LiffContext{SquareChat: &LineThrift.LiffSquareChatContext{SquareChatMid: to}}})
		if liffexcption, ok := err.(*LineThrift.LiffException); ok && liffexcption.Code == 3 {
			err = p.AllowLiffConsent(liffexcption.Payload.ConsentRequired.ConsentUrl)
			if err != nil {
				return err
			}
			token, err = p.LiffService().IssueLiffView(p.ctx, &LineThrift.LiffViewRequest{LiffId: LINE_LIFF_ID, Context: &LineThrift.LiffContext{SquareChat: &LineThrift.LiffSquareChatContext{SquareChatMid: to}}})
			if err != nil {
				return err
			}
		}
		accessToken = token.AccessToken
	} else {
		accessToken = token
	}
	client := &http.Client{}
	if p.Proxy != "" {
		proxyURL, _ := url.Parse(p.Proxy)
		client = &http.Client{Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		}}
	}
	url := LINE_API_DOMAIN + "/message/v3/share"
	datas := `{"messages": [` + data + `]}`
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte(datas)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+accessToken)
	res, err := client.Do(req)
	if err == nil {
		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)
		fmt.Println(string(body))
	}
	return err
}
