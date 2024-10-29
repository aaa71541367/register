package goline

import (
	"github.com/valyala/fasthttp"
)

func (p *LINE) getNewRequest(method, url string, body []byte, header map[string]string, cookies []*fasthttp.Cookie) (response *fasthttp.Response, err error) {
	request := fasthttp.AcquireRequest()
	response = fasthttp.AcquireResponse()
	request.SetRequestURI(url)
	request.SetBody(body)
	request.Header.SetMethod(method)
	request.Header.Set("User-Agent", p.userAgent)
	request.Header.Set("X-Line-Application", p.AppName)
	request.Header.Set("X-Line-Access", p.AuthToken)
	if header != nil {
		for key, value := range header {
			request.Header.Set(key, value)
		}
	}
	for _, cookie := range cookies {
		request.Header.SetCookie(string(cookie.Key()), string(cookie.Value()))
	}
	p.setDefaultHttpClient()
	err = p.fast_connection.DoRedirects(request, response, 5)
	return response, err
}

func (p *LINE) PostContent(url string, body []byte, header map[string]string) (res *fasthttp.Response, err error) {
	res, err = p.getNewRequest("POST", url, body, header, nil)
	return res, err
}

func (p *LINE) GetContent(url string, body []byte, header map[string]string) (res *fasthttp.Response, err error) {
	res, err = p.getNewRequest("GET", url, body, header, nil)
	return res, err
}

func (p *LINE) DeleteContent(url string, body []byte, header map[string]string) (res *fasthttp.Response, err error) {
	res, err = p.getNewRequest("DELETE", url, body, header, nil)
	return res, err
}

func (p *LINE) PutContent(url string, body []byte, header map[string]string) (res *fasthttp.Response, err error) {
	res, err = p.getNewRequest("PUT", url, body, header, nil)
	return res, err
}
