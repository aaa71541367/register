package goline

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"aaa71541367/register/LineThrift"
	"aaa71541367/register/thrift"

	"github.com/valyala/fasthttp"
)

var line_password = "mian89Ab"

type LINE struct {
	AuthToken       string              `json:"authToken"`
	Certificate     string              `json:"cert"`
	Revision        int64               `json:"revision"`
	MID             string              `json:"mid"`
	AppName         string              `json:"appName"`
	Number          int                 `json:"number"`
	Proxy           string              `json:"proxy"`
	E2EEKey         *LineThrift.E2EEKey `json:"e2eeKey"`
	GlobalRev       int64               `json:"globalRev"`
	IndividualRev   int64               `json:"individualRev"`
	Friends         []string            `json:"-"`
	Limit           bool                `json:"-"`
	pushClient      *LineThrift.TalkServiceClient
	connection      *http.Client
	fast_connection *fasthttp.Client
	ctx             context.Context
	Cancel          context.CancelFunc
	squareObsToken  string
	timeLineHeader  map[string]string
	liffTokens      map[string]string
	e2eePubKeys     map[string]*LineThrift.E2EEPublicKey
	reqSeq          int32
	userAgent       string
}

func NewLogin() *LINE {
	p := new(LINE)
	p.AuthToken = ""
	p.AppName = PRIMARY_DEVICE
	p.userAgent = "Line/" + strings.Split(p.AppName, "\t")[1]
	p.Certificate = ""
	p.Revision = 0
	p.MID = ""
	p.Limit = false
	p.timeLineHeader = map[string]string{}
	p.liffTokens = map[string]string{}
	p.e2eePubKeys = map[string]*LineThrift.E2EEPublicKey{}
	p.Proxy = ""
	p.reqSeq = 1000
	p.ctx, p.Cancel = context.WithCancel(context.TODO())
	p.Friends = []string{}
	return p
}

func (p *LINE) Login(args ...string) (err error) {
	if len(args) == 1 {
		p.AppName = PRIMARY_DEVICE
		p.userAgent = "Line/" + strings.Split(p.AppName, "\t")[1]
		p.MID = args[0][:33]
		err = p.LoginWithAuthToken(args[0])
	} else if len(args) == 2 {
		p.AppName = args[1]
		p.userAgent = "Line/" + strings.Split(p.AppName, "\t")[1]
		err = p.LoginWithAuthToken(args[0])
	} else {
		return fmt.Errorf("invalid param")
	}
	return err
}

type TalkServiceConnection struct {
	*LineThrift.TalkServiceClient
	Clinet *thrift.THttpClient
}

func PreloadConnection(token, appname string) TalkServiceConnection {
	httpClient, _ := thrift.NewTHttpClientWithOptions(LINE_HOST_DOMAIN+LINE_API_QUERY_PATH_SEC, thrift.THttpClientOptions{
		Client: &http.Client{
			Transport: http.DefaultTransport.(*http.Transport).Clone(),
		},
	})
	buffer := thrift.NewTBufferedTransportFactory(2048)
	trans := httpClient.(*thrift.THttpClient)
	trans.SetHeader("x-line-application", appname)
	trans.SetHeader("x-line-access", token)
	trans.SetHeader("User-Agent", "Line/"+strings.Split(appname, "\t")[1])
	trans.SetHeader("x-lal", "zh-Hant_TW")
	trans.SetHeader("x-lpv", "1")
	trans.SetHeader("Connection", "Keep-Alive")
	buftrans, _ := buffer.GetTransport(trans)
	compactProtocol := thrift.NewTCompactProtocolFactory()
	client := LineThrift.NewTalkServiceClientFactory(buftrans, compactProtocol)
	client.Noop(context.Background())
	return TalkServiceConnection{client, trans}
}
