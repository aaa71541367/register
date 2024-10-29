package goline

import (
	"github.com/shillbie/register/LineThrift"

	"github.com/valyala/fasthttp"
)

func (p *LINE) voteRequest(req *fasthttp.Request) (res *fasthttp.Response, err error) {
	resp, err := p.LiffService().IssueLiffView(p.ctx, &LineThrift.LiffViewRequest{LiffId: "1477715170-Pl2JnXpR", Context: &LineThrift.LiffContext{Chat: &LineThrift.LiffChatContext{ChatMid: "ccde797770ac0a163ced6e16b0baf2135"}}, DeviceSetting: &LineThrift.LiffDeviceSetting{VideoAutoPlayAllowed: true, AdvertisingId: &LineThrift.LiffAdvertisingId{AdvertisingId: "378c8582-9e2e-4273-a586-289eb6623087"}}})
	if err != nil {
		return res, err
	}
	req.Header.Set("User-Agent", p.userAgent)
	req.Header.Set("Connection", "Keep-Alive")
	req.Header.Set("X-Liff-Token", "Bearer "+resp.AccessToken)
	p.setDefaultHttpClient()
	res = fasthttp.AcquireResponse()
	err = p.fast_connection.Do(req, res)
	return res, err
}

func (p *LINE) GetVoteInfo(voteID, to string) (res string, err error) {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(LINE_W_DOMAIN + "/poll/ajax/poll/question/" + voteID)
	req.Header.SetMethod("GET")
	req.Header.Set("X-LINE-Chat-ID", to)
	resp, err := p.voteRequest(req)
	if err != nil {
		return res, err
	}
	return string(resp.Body()), nil
}

func (p *LINE) VoteChoise(voteID, to, choiseID string) (res string, err error) {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(LINE_W_DOMAIN + "/poll/ajax/poll/question/" + voteID + "/vote")
	req.SetBody([]byte("[\"" + choiseID + "\"]"))
	req.Header.SetMethod("POST")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-LINE-Chat-ID", to)
	resp, err := p.voteRequest(req)
	if err != nil {
		return res, err
	}
	return string(resp.Body()), nil
}
