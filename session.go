package goline

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"

	"github.com/shillbie/register/LineThrift"
	"github.com/shillbie/register/thrift"

	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpproxy"
	"golang.org/x/net/http2"
)

func (p *LINE) setDefaultHttpClient() {
	if p.fast_connection == nil {
		if p.Proxy != "" && !strings.HasPrefix(p.Proxy, "fake://") {
			fastproxy := strings.ReplaceAll(p.Proxy, "http://", "")
			fastproxy = strings.ReplaceAll(fastproxy, "http2://", "")
			p.fast_connection = &fasthttp.Client{Dial: fasthttpproxy.FasthttpHTTPDialer(fastproxy)}
		} else {
			p.fast_connection = &fasthttp.Client{}
		}
	}
	if p.connection == nil {
		if p.Proxy == "" || strings.HasPrefix(p.Proxy, "fake://") {
			trans := http.DefaultTransport.(*http.Transport).Clone()
			trans.MaxIdleConnsPerHost = 20
			p.connection = &http.Client{
				Transport: trans,
			}
		} else if proxy, err := url.Parse(p.Proxy); err == nil && strings.HasPrefix(p.Proxy, "http://") {
			p.connection = &http.Client{
				Transport: &http.Transport{
					Proxy: http.ProxyURL(proxy),
				},
			}
		} else if strings.HasPrefix(p.Proxy, "http2://") {
			http2proxy := strings.ReplaceAll(p.Proxy, "http2://", "")
			p.connection = &http.Client{
				Transport: &http2.Transport{
					AllowHTTP: true,
					DialTLS: func(network, addr string, cfg *tls.Config) (net.Conn, error) {
						// fmt.Println("new conn", network, addr)
						return net.Dial("tcp", http2proxy)
					},
				},
			}
		} else {
			panic("invalid proxy settings:" + p.Proxy)
		}
	}
}

func (p *LINE) TalkService() *LineThrift.TalkServiceClient {
	client := p.ThriftWrapper(LINE_API_QUERY_PATH_SEC, true)
	return LineThrift.NewTalkServiceClient(client)
}

func (p *LINE) SettingService() *LineThrift.TalkServiceClient {
	client := p.ThriftWrapper(LINE_API_QUERY_PATH_SEC, true)
	return LineThrift.NewTalkServiceClient(client)
}

func (p *LINE) RelationService() *LineThrift.RelationServiceClient {
	client := p.ThriftWrapper(LINE_RELATION_QUERY_PATH, true)
	return LineThrift.NewRelationServiceClient(client)
}

func (p *LINE) PollService() *LineThrift.TalkServiceClient {
	client := p.ThriftWrapper(LINE_POLL_QUERY_PATH_FIR, true)
	return LineThrift.NewTalkServiceClient(client)
}

func (p *LINE) ChannelService() *LineThrift.ChannelServiceClient {
	client := p.ThriftWrapper(LINE_CHAN_QUERY_PATH, true)
	return LineThrift.NewChannelServiceClient(client)
}

func (p *LINE) CallService() *LineThrift.CallServiceClient {
	client := p.ThriftWrapper(LINE_CALL_QUERY_PATH, true)
	return LineThrift.NewCallServiceClient(client)
}

func (p *LINE) SquareService() *LineThrift.SquareServiceClient {
	client := p.ThriftWrapper(LINE_SQUARE_QUERY_PATH, true)
	return LineThrift.NewSquareServiceClient(client)
}

func (p *LINE) LiffService() *LineThrift.LiffServiceClient {
	client := p.ThriftWrapper(LINE_LIFF_QUERY_PATH, true)
	return LineThrift.NewLiffServiceClient(client)
}
func (p *LINE) AuthService() *LineThrift.TalkServiceClient {
	client := p.ThriftWrapper(LINE_LOGIN_QUERY_PATH, false)
	return LineThrift.NewTalkServiceClient(client)
}

func (p *LINE) VerifyService() *LineThrift.AuthServiceClient {
	client := p.ThriftWrapper(LINE_AUTH_QUERY_PATH, true)
	return LineThrift.NewAuthServiceClient(client)
}

func (p *LINE) LoginZService() *LineThrift.AuthServiceClient {
	client := p.ThriftWrapper(LINE_AUTH_LOGIN_QUERY_PATH, true)
	return LineThrift.NewAuthServiceClient(client)
}

func (p *LINE) RefreshTokenService() *LineThrift.AccessTokenRefreshServiceClient {
	client := p.ThriftWrapper("/EXT/auth/tokenrefresh/v1", true)
	return LineThrift.NewAccessTokenRefreshServiceClient(client)
}

func (p *LINE) QrService() *LineThrift.SecondaryQrCodeLoginServiceClient {
	client := p.ThriftWrapper(LINE_SQ_QUERY_PATH, false)
	return LineThrift.NewSecondaryQrCodeLoginServiceClient(client)
}

func (p *LINE) QrVService() *LineThrift.SecondaryQrCodeLoginServiceClient {
	client := p.ThriftWrapper(LINE_SQ_VERIFY_QUERY_PATH, true)
	return LineThrift.NewSecondaryQrCodeLoginServiceClient(client)
}

func (p *LINE) PwlessService() *LineThrift.SecondaryPwlessLoginServiceClient {
	client := p.ThriftWrapper("/acct/lgn/secpwless/v1", false)
	return LineThrift.NewSecondaryPwlessLoginServiceClient(client)
}

func (p *LINE) PwlessVerifyService() *LineThrift.SecondaryPwlessLoginPermitNoticeServiceClient {
	client := p.ThriftWrapper("/acct/lp/lgn/secpwless/v1", true)
	return LineThrift.NewSecondaryPwlessLoginPermitNoticeServiceClient(client)
}

func (p *LINE) PwlessPinService() *LineThrift.SecondAuthFactorPinCodeServiceClient {
	client := p.ThriftWrapper("/ACCT/authfactor/second/pincode/v1", false)
	return LineThrift.NewSecondAuthFactorPinCodeServiceClient(client)
}

func (p *LINE) PwlessLoginService() *LineThrift.SecondaryPwlessLoginPermitServiceClient {
	client := p.ThriftWrapper("/ACCT/lgn/secpwless/v1", true)
	return LineThrift.NewSecondaryPwlessLoginPermitServiceClient(client)
}

func (p *LINE) SnsService() *LineThrift.SnsAdaptorServiceClient {
	client := p.ThriftWrapper("/api/v4p/sa", false)
	return LineThrift.NewSnsAdaptorServiceClient(client)
}

func (p *LINE) PrimaryQRMigrationLPService() *LineThrift.PrimaryQrCodeMigrationLongPollingServiceClient {
	client := p.ThriftWrapper("/EXT/auth/feature-user/lp/api/primary/mig/qr", true)
	return LineThrift.NewPrimaryQrCodeMigrationLongPollingServiceClient(client)
}

func (p *LINE) PrimaryQRMigrationPreService() *LineThrift.PrimaryQrCodeMigrationPreparationServiceClient {
	client := p.ThriftWrapper("/EXT/auth/feature-user/api/primary/mig/qr/prepare", true)
	return LineThrift.NewPrimaryQrCodeMigrationPreparationServiceClient(client)
}

func (p *LINE) PrimaryQRMigrationService() *LineThrift.PrimaryQrCodeMigrationServiceClient {
	client := p.ThriftWrapper("/ext/auth/feature-guest/api/primary/mig/qr", true)
	return LineThrift.NewPrimaryQrCodeMigrationServiceClient(client)
}

func (p *LINE) ProxyService() *LineThrift.ProxyServiceClient {
	client := p.ThriftWrapper("/proxy", true)
	return LineThrift.NewProxyServiceClient(client)
}

func (p *LINE) ThriftWrapper(apiUrl string, withToken bool) *thrift.TStandardClient {
	p.setDefaultHttpClient()
	var httpClient thrift.TTransport
	var protocol thrift.TProtocol
	httpClient, _ = thrift.NewTHttpClientWithOptions(LINE_HOST_DOMAIN+apiUrl, thrift.THttpClientOptions{
		Client: p.connection,
	})
	buffer := thrift.NewTBufferedTransportFactory(2048)
	trans := httpClient.(*thrift.THttpClient)
	if withToken && p.AuthToken != "" {
		trans.SetHeader("x-line-access", p.AuthToken)
	}
	if strings.HasPrefix(p.Proxy, "fake://") {
		ip, err := url.Parse(p.Proxy)
		if err == nil {
			trans.SetHeader("x-forwarded-for", ip.Hostname())
		} else {
			fmt.Println("NO IP FOUND")
		}
	}
	if apiUrl == LINE_POLL_QUERY_PATH_FIR || apiUrl == LINE_POLL_QUERY_PATH_SEC || apiUrl == LINE_POLL_QUERY_PATH_THI {
		trans.SetHeader("x-las", "F")
		trans.SetHeader("x-lam", "w")
		trans.SetHeader("x-lac", "44125")
	}
	trans.SetHeader("User-Agent", p.userAgent)
	trans.SetHeader("X-Line-Application", p.AppName)
	trans.SetHeader("x-lal", "zh-Hant_TW")
	trans.SetHeader("x-lpv", "1")
	trans.SetHeader("Connection", "Keep-Alive")
	buftrans, _ := buffer.GetTransport(trans)
	protocol = thrift.NewTCompactProtocolFactory().GetProtocol(buftrans)
	thriftClient := thrift.NewTStandardClient(protocol, protocol)
	return thriftClient
}

func (p *LINE) FastThriftWrapper(apiUrl string, withToken bool) *thrift.TStandardClient {
	p.setDefaultHttpClient()
	var httpClient thrift.TTransport
	var protocol thrift.TProtocol
	httpClient, _ = thrift.NewTFastHttpClientWithOptions(LINE_HOST_DOMAIN+apiUrl, thrift.TFastHttpClientOptions{
		Client: p.fast_connection,
	})
	buffer := thrift.NewTBufferedTransportFactory(2048)
	trans := httpClient.(*thrift.TFastHttpClient)
	if withToken && p.AuthToken != "" {
		trans.SetHeader("x-line-access", p.AuthToken)
	}
	if strings.HasPrefix(p.Proxy, "fake://") {
		ip, err := url.Parse(p.Proxy)
		if err == nil {
			trans.SetHeader("x-forwarded-for", ip.Hostname())
		} else {
			fmt.Println("NO IP FOUND")
		}
	}
	if apiUrl == LINE_POLL_QUERY_PATH_FIR || apiUrl == LINE_POLL_QUERY_PATH_SEC || apiUrl == LINE_POLL_QUERY_PATH_THI {
		trans.SetHeader("x-las", "F")
		trans.SetHeader("x-lam", "w")
		trans.SetHeader("x-lac", "44125")
	}
	trans.SetHeader("User-Agent", p.userAgent)
	trans.SetHeader("X-Line-Application", p.AppName)
	trans.SetHeader("x-lal", "zh-Hant_TW")
	trans.SetHeader("x-lpv", "1")
	trans.SetHeader("Connection", "Keep-Alive")
	buftrans, _ := buffer.GetTransport(trans)
	protocol = thrift.NewTCompactProtocolFactory().GetProtocol(buftrans)
	thriftClient := thrift.NewTStandardClient(protocol, protocol)
	return thriftClient
}
