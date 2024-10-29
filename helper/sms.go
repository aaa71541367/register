package helper

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func GetSimHni(country string) string {
	if data, ok := countryInfo[country]; ok {
		return data.Hni
	}
	return "46692"
}

func GetSimCarrier(country string) string {
	if data, ok := countryInfo[country]; ok {
		return data.CarrierName
	}
	return "Chunghwa"
}
func GetRandomIP(country string) (net.IP, error) {
	if data, ok := countryInfo[country]; ok && len(data.CIDRs) > 0 {
		rand.Seed(time.Now().Unix())
		cidr := data.CIDRs[rand.Intn(len(data.CIDRs))]
	GENERATE:
		ip, ipnet, err := net.ParseCIDR(cidr)
		if err != nil {
			return nil, err
		}
		// The number of leading 1s in the mask
		ones, _ := ipnet.Mask.Size()
		quotient := ones / 8
		remainder := ones % 8
		// create random 4-byte byte slice
		r := make([]byte, 4)
		rand.Read(r)
		for i := 0; i <= quotient; i++ {
			if i == quotient {
				shifted := byte(r[i]) >> remainder
				r[i] = ipnet.IP[i] + shifted
			} else {
				r[i] = ipnet.IP[i]
			}
		}
		ip = net.IPv4(r[0], r[1], r[2], r[3])
		if ip.Equal(ipnet.IP) /*|| ip.Equal(broadcast) */ {
			// we got unlucky. The host portion of our ipv4 address was
			// either all 0s (the network address) or all 1s (the broadcast address)
			goto GENERATE
		}
		return ip, nil
	}
	return nil, fmt.Errorf("no cidr provide")
}

func removeCountryCode(phone, country string) string {
	if strings.HasPrefix(phone, "+") {
		phone = phone[1:]
	}
	if code, ok := countryInfo[strings.ToLower(country)]; ok && strings.HasPrefix(phone, code.CountryCode) {
		phone = phone[len(code.CountryCode):]
	}
	return phone
}

type SmsReciever interface {
	GetPhoneNum(country string) (phone, id string, err error)
	GetPhoneCode(id string, times int) (code string, err error)
	ReleasePhone(id, act string) (err error)
}

type FiveSim struct {
	apiToken  string
	countryId map[string]string
	status    map[string]string
	client    *http.Client
}

func NewFiveSim(apiToken string) SmsReciever {
	return &FiveSim{
		apiToken: apiToken,
		countryId: map[string]string{
			"lt": "lithuania",
			"tw": "taiwan",
			"ca": "canada",
			"us": "usa",
			"hk": "hongkong",
			"no": "norway",
			"dk": "denmark",
			"fi": "finland",
			"id": "indonesia",
			"th": "thailand",
			"ru": "russia",
		},
		status: map[string]string{
			"shieldNumber":  "cancel",
			"releaseNumber": "cancel",
		},
		client: new(http.Client),
	}
}

func (p *FiveSim) GetPhoneNum(country string) (phone, id string, err error) {
	if code, ok := p.countryId[country]; ok {
		req, _ := http.NewRequest("GET", "https://5sim.net/v1/user/buy/activation/"+code+"/any/line", nil)
		req.Header.Set("Authorization", "Bearer "+p.apiToken)
		req.Header.Set("Accept", "application/json")
		res, err := p.client.Do(req)
		if err != nil {
			return "", "", err
		}
		defer res.Body.Close()
		bData, _ := ioutil.ReadAll(res.Body)
		result := struct {
			ID      int    `json:"id"`
			Phone   string `json:"phone"`
			Country string `json:"country"`
			Status  string `json:"status"`
			Sms     []struct {
				Sender string `json:"sender"`
				Code   string `json:"code"`
				Text   string `json:"text"`
			} `json:"sms"`
		}{}
		err = json.Unmarshal(bData, &result)
		if err != nil || len(result.Phone) < 5 {
			return "", "", fmt.Errorf(string(bData))
		}
		result.Phone = removeCountryCode(result.Phone, country)
		return result.Phone, fmt.Sprintf("%d", result.ID), nil
	}
	return "", "", fmt.Errorf("error country")
}

func (p *FiveSim) GetPhoneCode(id string, times int) (code string, err error) {
	defer fmt.Println("")
	for i := 0; i < times; i++ {
		req, _ := http.NewRequest("GET", "https://5sim.net/v1/user/check/"+id, nil)
		req.Header.Set("Authorization", "Bearer "+p.apiToken)
		req.Header.Set("Accept", "application/json")
		res, err := p.client.Do(req)
		if err != nil {
			return "", err
		}
		defer res.Body.Close()
		bData, _ := ioutil.ReadAll(res.Body)
		result := struct {
			ID      int    `json:"id"`
			Phone   string `json:"phone"`
			Country string `json:"country"`
			Status  string `json:"status"`
			Sms     []struct {
				Sender string `json:"sender"`
				Code   string `json:"code"`
				Text   string `json:"text"`
			} `json:"sms"`
		}{}
		err = json.Unmarshal(bData, &result)
		if err != nil {
			return "", err
		} else if result.Status == "PENDING" {
			fmt.Print("\r正在等待驗證碼...", i+1, "...")
			time.Sleep(2 * time.Second)
			continue
		} else if result.Status == "RECEIVED" {
			if len(result.Sms) == 1 {
				return result.Sms[0].Code, nil
			} else if len(result.Sms) == 0 {
				fmt.Print("\r正在等待驗證碼...", i+1, "...")
				time.Sleep(2 * time.Second)
				continue
			} else {
				fmt.Println(result)
				return "", fmt.Errorf("Some error.")
			}
		} else {
			return "", fmt.Errorf(result.Status)
		}
	}
	return "", fmt.Errorf("Timeout.")
}

func (p *FiveSim) ReleasePhone(id, act string) (err error) {
	req, _ := http.NewRequest("GET", "https://5sim.net/v1/user/"+p.status[act]+"/"+id, nil)
	req.Header.Set("Authorization", "Bearer "+p.apiToken)
	req.Header.Set("Accept", "application/json")
	_, err = p.client.Do(req)
	return err
}

type OnlineSim struct {
	apiToken  string
	countryId map[string]string
	status    map[string]string
	client    *http.Client
}

func NewOnlineSim(apiToken string) SmsReciever {
	tls_config := &tls.Config{}
	tls_config.InsecureSkipVerify = true
	return &OnlineSim{
		apiToken: apiToken,
		countryId: map[string]string{
			"ru": "7",
		},
		status: map[string]string{
			"shieldNumber":  "cancel",
			"releaseNumber": "cancel",
		},
		client: &http.Client{Transport: &http.Transport{TLSClientConfig: tls_config}},
	}
}

func (p *OnlineSim) GetPhoneNum(country string) (phone, id string, err error) {
	if code, ok := p.countryId[country]; ok {
		param := url.Values{}
		param.Set("apikey", p.apiToken)
		param.Set("country", code)
		param.Set("service", "LineMessenger")
		param.Set("lang", "en")
		req, _ := http.NewRequest("GET", "https://onlinesim.ru/api/getNum.php?"+param.Encode(), nil)
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36")
		res, err := p.client.Do(req)
		if err != nil {
			return "", "", err
		}
		bData, _ := ioutil.ReadAll(res.Body)
		res.Body.Close()
		result := struct {
			Response int `json:"response"`
			Tzid     int `json:"tzid"`
		}{}
		err = json.Unmarshal(bData, &result)
		if err != nil || result.Response != 1 {
			return "", "", fmt.Errorf(string(bData))
		}
		param.Set("tzid", fmt.Sprintf("%d", result.Tzid))
		req, _ = http.NewRequest("GET", "https://onlinesim.ru/api/getState.php?"+param.Encode(), nil)
		req.Header.Set("Accept", "application/json")
		res, err = p.client.Do(req)
		if err != nil {
			return "", "", err
		}
		bData, _ = ioutil.ReadAll(res.Body)
		res.Body.Close()
		result2 := []struct {
			Response string `json:"response"`
			Tzid     int    `json:"tzid"`
			Service  string `json:"service"`
			Number   string `json:"number"`
			Message  string `json:"msg"`
		}{}
		json.Unmarshal(bData, &result2)
		result2[0].Number = removeCountryCode(result2[0].Number, country)
		return result2[0].Number, fmt.Sprintf("%d", result.Tzid), nil
	}
	return "", "", fmt.Errorf("error country")
}

func (p *OnlineSim) GetPhoneCode(id string, times int) (code string, err error) {
	defer fmt.Println("")
	param := url.Values{}
	param.Set("apikey", p.apiToken)
	param.Set("tzid", id)
	for i := 0; i < times; i++ {
		req, _ := http.NewRequest("GET", "https://onlinesim.ru/api/getState.php?"+param.Encode(), nil)
		req.Header.Set("Accept", "application/json")
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36")
		req.Header.Set("Accept-Language", "en-us")
		res, err := p.client.Do(req)
		if err != nil {
			return "", err
		}
		defer res.Body.Close()
		bData, _ := ioutil.ReadAll(res.Body)
		result := []struct {
			Response string `json:"response"`
			Tzid     int    `json:"tzid"`
			Service  string `json:"service"`
			Number   string `json:"number"`
			Message  string `json:"msg"`
		}{}
		json.Unmarshal(bData, &result)
		if result[0].Response == "TZ_NUM_ANSWER" {
			return result[0].Message, nil
		} else if result[0].Response == "TZ_NUM_WAIT" {
			fmt.Print("\r正在等待驗證碼...", i+1, "...")
			time.Sleep(2 * time.Second)
			continue
		} else {
			return "", fmt.Errorf(result[0].Response)
		}
	}
	return "", fmt.Errorf("Timeout.")
}

func (p *OnlineSim) ReleasePhone(id, act string) (err error) {
	param := url.Values{}
	param.Set("apikey", p.apiToken)
	param.Set("tzid", id)
	req, _ := http.NewRequest("GET", "https://onlinesim.ru/api/setOperationOk.php?"+param.Encode(), nil)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36")
	req.Header.Set("Accept-Language", "en-us")
	_, err = p.client.Do(req)
	return err
}

type SmsManual struct{}

func NewSmsManual() SmsReciever {
	return &SmsManual{}
}

func (p *SmsManual) GetPhoneNum(country string) (phone, id string, err error) {
	fmt.Print("PHONE:")
	fmt.Scan(&phone)
	phone = removeCountryCode(phone, country)
	return phone, "", nil
}

func (p *SmsManual) GetPhoneCode(id string, times int) (code string, err error) {
	fmt.Print("CODE:")
	fmt.Scan(&code)
	return code, nil
}

func (p *SmsManual) ReleasePhone(id, act string) error {
	return nil
}

type SmsActivate struct {
	apiToken  string
	countryId map[string]int
	status    map[string]int
}

func NewSmsActivate(apiToken string) SmsReciever {
	return &SmsActivate{
		apiToken: "7818A1bbb795d0b5eb73d8069637e40A",
		countryId: map[string]int{
			"id":  6,
			"hk":  14,
			"my":  7,
			"in":  22,
			"vn":  10,
			"lt":  44,
			"ca":  36,
			"phi": 4,
			"ru":  0,
			"us":  187,
			"kg":  11,
			"th":  52,
			"tr":  62,
			"es":  56,
			"gy":  131,
		},
		status: map[string]int{
			"shieldNumber":  8,
			"releaseNumber": 6,
		},
	}
}

func (p *SmsActivate) GetPhoneNum(country string) (phone, id string, err error) {
	if id, ok := p.countryId[country]; ok {
		res, err := http.Get(fmt.Sprintf("https://sms-activate.org/stubs/handler_api.php?api_key=7818A1bbb795d0b5eb73d8069637e40A&action=getNumber&service=me&forward=0&country=%d&verification=false", id))
		if err != nil {
			return "", "", err
		}
		defer res.Body.Close()
		dataresp, _ := ioutil.ReadAll(res.Body)
		if datas, ok := FormatCommand(string(dataresp), "([0-9]{1,}):([0-9]{1,})", "ACCESS_NUMBER:"); ok {
			if len(datas) == 2 {
				datas[1] = removeCountryCode(datas[1], country)
				return datas[1], datas[0], nil
			}
		}
		return "", "", fmt.Errorf(string(dataresp))
	}
	return "", "", fmt.Errorf("error country")
}

func (p *SmsActivate) GetPhoneCode(id string, times int) (code string, err error) {
	defer fmt.Println("")
	for i := 0; i < times; i++ {
		res, err := http.Get(fmt.Sprintf("https://api.sms-activate.org/stubs/handler_api.php?api_key=7818A1bbb795d0b5eb73d8069637e40A&action=getStatus&id=%s", id))
		if err != nil {
			continue
		}
		defer res.Body.Close()
		dataresp, _ := ioutil.ReadAll(res.Body)
		if strings.HasPrefix(string(dataresp), "STATUS_WAIT_CODE") {
			fmt.Print("\r正在等待驗證碼...", i+1, "...")
			time.Sleep(2 * time.Second)
			continue
		} else if strings.HasPrefix(string(dataresp), "STATUS_OK:") {
			return string(dataresp)[10:], nil
		} else {
			return "", fmt.Errorf(string(dataresp))
		}
	}
	return "", fmt.Errorf("Timeout.")
}

func (p *SmsActivate) ReleasePhone(id, act string) (err error) {
	_, err = http.Get(fmt.Sprintf("https://api.sms-activate.org/stubs/handler_api.php?api_key=7818A1bbb795d0b5eb73d8069637e40A&action=setStatus&id=%s&status=%d&forward=0",id, p.status[act]))
	return err
}

type Smsnpng struct {
	apiKey      string
	countryCode map[string]int
}

func NewSmsnpng(apiKey string) SmsReciever {
	return &Smsnpng{
		apiKey: apiKey,
		countryCode: map[string]int{
			// "us": 162,
			"my": 163,
			"th": 164,
			"id": 166,
			"hk": 165,
			// US: 109 MY: 112 TH: 128 ID: 113
		},
	}
}

func (p *Smsnpng) GetPhoneNum(country string) (phone, id string, err error) {
	if id, ok := p.countryCode[country]; ok {
		res, err := http.Get("http://www.smsnpng.com/api/getNumber?api_token=" + p.apiKey + "&platform=5&projectId=496&countryId=" + fmt.Sprintf("%d", id))
		if err != nil {
			return "", "", err
		}
		defer res.Body.Close()
		sitemap, _ := ioutil.ReadAll(res.Body)
		result := struct {
			Code    int    `json:"code"`
			Message string `json:"msg"`
			Data    struct {
				Number string `json:"number"`
				ID     int    `json:"id"`
			} `json:"data"`
		}{}
		err = json.Unmarshal(sitemap, &result)
		if err != nil || result.Code != 1 || len(result.Data.Number) < 5 {
			return "", "", fmt.Errorf(result.Message)
		}
		result.Data.Number = removeCountryCode(result.Data.Number, country)
		return result.Data.Number, fmt.Sprintf("%d", result.Data.ID), nil
	}
	return "", "", fmt.Errorf("error country")
}

func (p *Smsnpng) GetPhoneCode(id string, times int) (code string, err error) {
	defer fmt.Println("")
	for i := 0; i < times; i++ {
		res, err := http.Get("http://www.smsnpng.com/api/getSmsContent?api_token=" + p.apiKey + "&orderid=" + id)
		if err != nil {
			continue
		}
		defer res.Body.Close()
		sitemap, _ := ioutil.ReadAll(res.Body)
		result := struct {
			Code    int    `json:"code"`
			Message string `json:"msg"`
			Data    struct {
				Times    string `json:"times"`
				Messages string `json:"messages"`
			} `json:"data"`
		}{}
		err = json.Unmarshal(sitemap, &result)
		if err != nil || result.Code != 1 {
			fmt.Print("\r正在等待驗證碼...", i+1, "...")
			time.Sleep(1 * time.Second)
			continue
		} else {
			return result.Data.Messages, nil
		}
	}
	return "", fmt.Errorf("Timeout.")
}

func (p *Smsnpng) ReleasePhone(id, act string) error {
	_, err := http.Get("http://www.smsnpng.com/api/" + act + "?api_token=" + p.apiKey + "&orderid=" + id)
	return err
}

var countryInfo = map[string]struct {
	Hni         string
	CarrierName string
	CountryCode string
	CIDRs       []string
}{
	"id": {
		Hni:         "51001",
		CarrierName: "INDOSAT",
		CountryCode: "62",
		CIDRs: []string{
			"202.155.0.0/22",
			"202.155.100.0/23",
			"202.155.104.0/23",
			"202.155.106.0/23",
			"202.155.110.0/23",
			"202.155.112.0/23",
			"114.10.10.0/23",
			"114.120.0.0/22",
			"114.120.4.0/22",
			"114.120.8.0/22",
			"114.120.12.0/22",
			"114.120.16.0/22",
			"114.120.20.0/22",
			"114.120.24.0/22",
			"114.120.28.0/22",
			"114.120.32.0/22",
			"114.120.36.0/22",
			"114.120.40.0/22",
			"114.120.44.0/22",
			"114.120.48.0/22",
			"114.120.52.0/22",
			"114.120.56.0/22",
			"114.120.60.0/22",
			"114.120.64.0/22",
			"114.120.68.0/22",
			"114.120.72.0/22",
			"114.120.76.0/22",
			"114.120.80.0/22",
			"114.120.84.0/22",
			"114.120.88.0/22",
			"114.120.92.0/22",
			"114.120.96.0/22",
			"114.120.100.0/22",
			"114.120.104.0/22",
			"114.120.108.0/22",
			"114.120.112.0/22",
			"114.120.116.0/22",
			"114.120.120.0/22",
			"114.120.124.0/22",
			"114.120.128.0/22",
			"114.120.132.0/22",
			"114.120.136.0/22",
			"114.120.140.0/22",
			"114.120.144.0/22",
			"114.120.148.0/22",
			"114.120.152.0/22",
			"114.120.156.0/22",
			"114.120.160.0/22",
			"114.120.164.0/22",
			"114.120.168.0/22",
			"114.120.172.0/22",
			"114.120.176.0/22",
			"114.120.180.0/22",
			"114.120.184.0/22",
			"114.120.188.0/22",
			"114.120.192.0/22",
			"114.120.196.0/22",
			"114.120.200.0/22",
			"114.120.204.0/22",
			"114.120.208.0/22",
			"114.120.212.0/22",
			"114.120.216.0/22",
			"114.120.220.0/22",
			"114.120.224.0/22",
			"114.120.228.0/22",
			"114.120.232.0/22",
			"114.120.236.0/22",
			"114.120.240.0/22",
			"114.120.244.0/22",
			"114.120.248.0/22",
			"114.120.252.0/22",
			"114.121.0.0/22",
			"114.121.4.0/22",
			"114.121.8.0/22",
			"114.121.12.0/22",
			"114.121.16.0/22",
			"114.121.20.0/22",
			"114.121.24.0/22",
			"114.121.28.0/22",
			"114.121.32.0/22",
			"114.121.36.0/22",
			"114.121.40.0/22",
			"114.121.44.0/22",
			"114.121.48.0/22",
			"114.121.52.0/22",
			"114.121.56.0/22",
			"114.121.60.0/22",
			"114.121.64.0/22",
			"114.121.68.0/22",
			"114.121.72.0/22",
			"114.121.76.0/22",
			"114.121.80.0/22",
			"114.121.84.0/22",
			"114.121.88.0/22",
			"114.121.92.0/22",
			"114.121.96.0/22",
			"114.121.100.0/22",
			"114.121.104.0/22",
			"114.121.108.0/22",
			"114.121.112.0/22",
			"114.121.116.0/22",
			"114.121.120.0/22",
			"114.121.124.0/22",
			"114.121.128.0/22",
			"114.121.132.0/22",
			"114.121.136.0/22",
			"114.121.140.0/22",
			"114.121.144.0/22",
			"114.121.148.0/22",
			"114.121.152.0/22",
			"114.121.156.0/22",
			"114.121.160.0/22",
			"114.121.164.0/22",
			"114.121.168.0/22",
			"114.121.172.0/22",
			"114.121.176.0/22",
			"114.121.180.0/22",
			"114.121.184.0/22",
			"114.121.188.0/22",
			"114.121.192.0/22",
			"114.121.196.0/22",
			"114.121.200.0/22",
			"114.121.204.0/22",
			"114.121.208.0/22",
			"114.121.212.0/22",
			"114.121.216.0/22",
			"114.121.220.0/22",
			"114.121.224.0/22",
			"114.121.228.0/22",
			"114.121.232.0/22",
			"114.121.236.0/22",
			"114.121.240.0/22",
			"114.121.244.0/22",
			"114.121.248.0/22",
			"114.121.252.0/22",
			"114.122.0.0/22",
			"114.122.4.0/22",
			"114.122.8.0/22",
			"114.122.12.0/22",
			"114.122.16.0/22",
			"114.122.20.0/22",
			"114.122.24.0/22",
			"114.122.28.0/22",
			"114.122.32.0/22",
			"114.122.36.0/22",
			"114.122.40.0/22",
			"114.122.44.0/22",
			"114.122.48.0/22",
			"114.122.52.0/22",
			"114.122.56.0/22",
			"114.122.60.0/22",
			"114.122.64.0/22",
			"114.122.68.0/22",
			"114.122.72.0/22",
			"114.122.76.0/22",
			"114.122.80.0/22",
			"114.122.84.0/22",
			"114.122.88.0/22",
			"114.122.92.0/22",
			"114.122.96.0/22",
			"114.122.100.0/22",
			"114.122.104.0/22",
			"114.122.108.0/22",
			"114.122.112.0/22",
			"114.122.116.0/22",
			"114.122.120.0/22",
			"114.122.124.0/22",
			"114.122.128.0/22",
			"114.122.132.0/22",
			"114.122.136.0/22",
			"114.122.140.0/22",
			"114.122.144.0/22",
			"114.122.148.0/22",
			"114.122.152.0/22",
			"114.122.156.0/22",
			"114.122.160.0/22",
			"114.122.164.0/22",
			"114.122.168.0/22",
			"114.122.172.0/22",
			"114.122.176.0/22",
			"114.122.180.0/22",
			"114.122.184.0/22",
			"114.122.188.0/22",
			"114.122.192.0/22",
			"114.122.196.0/22",
			"114.122.200.0/22",
			"114.122.204.0/22",
			"114.122.208.0/22",
			"114.122.212.0/22",
			"114.122.216.0/22",
			"114.122.220.0/22",
			"114.122.224.0/22",
			"114.122.228.0/22",
			"114.122.232.0/22",
			"114.122.236.0/22",
			"114.122.240.0/22",
			"114.122.244.0/22",
			"114.122.248.0/22",
			"114.122.252.0/22",
			"114.123.0.0/22",
			"114.123.4.0/22",
			"114.123.8.0/22",
			"114.123.12.0/22",
			"114.123.16.0/22",
			"114.123.20.0/22",
			"114.123.24.0/22",
			"114.123.28.0/22",
			"114.123.32.0/22",
			"114.123.36.0/22",
			"114.123.40.0/22",
			"114.123.44.0/22",
			"114.123.48.0/22",
			"114.123.52.0/22",
			"114.123.56.0/22",
			"114.123.60.0/22",
			"114.123.64.0/22",
			"114.123.68.0/22",
			"114.123.72.0/22",
			"114.123.76.0/22",
			"114.123.80.0/22",
			"114.123.84.0/22",
			"114.123.88.0/22",
			"114.123.92.0/22",
			"114.123.96.0/22",
			"114.123.100.0/22",
			"114.123.104.0/22",
			"114.123.108.0/22",
			"114.123.112.0/22",
			"114.123.116.0/22",
			"114.123.120.0/22",
			"114.123.124.0/22",
			"114.123.128.0/22",
			"114.123.132.0/22",
			"114.123.136.0/22",
			"114.123.140.0/22",
			"114.123.144.0/22",
			"114.123.148.0/22",
			"114.123.152.0/22",
			"114.123.156.0/22",
			"114.123.160.0/22",
			"114.123.164.0/22",
			"114.123.168.0/22",
			"114.123.172.0/22",
			"114.123.176.0/22",
			"114.123.180.0/22",
			"114.123.184.0/22",
			"114.123.188.0/22",
			"114.123.192.0/22",
			"114.123.196.0/22",
			"114.123.200.0/22",
			"114.123.204.0/22",
			"114.123.208.0/22",
			"114.123.212.0/22",
			"114.123.216.0/22",
			"114.123.220.0/22",
			"114.123.224.0/22",
			"114.123.228.0/22",
			"114.123.232.0/22",
			"114.123.236.0/22",
			"114.123.240.0/22",
			"114.123.244.0/22",
			"114.123.248.0/22",
			"114.123.252.0/22",
			"114.124.0.0/22",
			"114.124.4.0/22",
			"114.124.8.0/22",
			"114.124.12.0/22",
			"114.124.16.0/22",
			"114.124.20.0/22",
			"114.124.24.0/22",
			"114.124.28.0/22",
			"114.124.32.0/22",
			"114.124.36.0/22",
			"114.124.40.0/22",
			"114.124.44.0/22",
			"114.124.48.0/22",
			"114.124.52.0/22",
			"114.124.56.0/22",
			"114.124.60.0/22",
			"114.124.64.0/22",
			"114.124.68.0/22",
			"114.124.72.0/22",
			"114.124.76.0/22",
			"114.124.80.0/22",
			"114.124.84.0/22",
			"114.124.88.0/22",
			"114.124.92.0/22",
			"114.124.96.0/22",
			"114.124.100.0/22",
			"114.124.104.0/22",
			"114.124.108.0/22",
			"114.124.112.0/22",
			"114.124.116.0/22",
			"114.124.120.0/22",
			"114.124.124.0/22",
			"114.124.128.0/22",
			"114.124.132.0/22",
			"114.124.136.0/22",
			"114.124.140.0/22",
			"114.124.144.0/22",
			"114.124.148.0/22",
			"114.124.152.0/22",
			"114.124.156.0/22",
			"114.124.160.0/22",
			"114.124.164.0/22",
			"114.124.168.0/22",
			"114.124.172.0/22",
			"114.124.176.0/22",
			"114.124.180.0/22",
			"114.124.184.0/22",
			"114.124.188.0/22",
			"114.124.192.0/22",
			"114.124.196.0/22",
			"114.124.200.0/22",
			"114.124.204.0/22",
			"114.124.208.0/22",
			"114.124.212.0/22",
			"114.124.216.0/22",
			"114.124.220.0/22",
			"114.124.224.0/22",
			"114.124.228.0/22",
			"114.124.232.0/22",
			"114.124.236.0/22",
			"114.124.240.0/22",
			"114.124.244.0/22",
			"114.124.248.0/22",
			"114.124.252.0/22",
			"114.125.0.0/22",
			"114.125.4.0/22",
			"114.125.8.0/22",
			"114.125.12.0/22",
			"114.125.16.0/22",
			"114.125.20.0/22",
			"114.125.24.0/22",
			"114.125.28.0/22",
			"114.125.32.0/22",
			"114.125.36.0/22",
			"114.125.40.0/22",
			"114.125.44.0/22",
			"114.125.48.0/22",
			"114.125.52.0/22",
			"114.125.56.0/22",
			"114.125.60.0/22",
			"114.125.64.0/22",
			"114.125.68.0/22",
			"114.125.72.0/22",
			"114.125.76.0/22",
			"114.125.80.0/22",
			"114.125.84.0/22",
			"114.125.88.0/22",
			"114.125.92.0/22",
			"114.125.96.0/22",
			"114.125.100.0/22",
			"114.125.104.0/22",
			"114.125.108.0/22",
			"114.125.112.0/22",
			"114.125.116.0/22",
			"114.125.120.0/22",
			"114.125.124.0/22",
			"114.125.128.0/22",
			"114.125.132.0/22",
			"114.125.136.0/22",
			"114.125.140.0/22",
			"114.125.144.0/22",
			"114.125.148.0/22",
			"114.125.152.0/22",
			"114.125.156.0/22",
			"114.125.160.0/22",
			"114.125.164.0/22",
			"114.125.168.0/22",
			"114.125.172.0/22",
			"114.125.176.0/22",
			"114.125.180.0/22",
			"114.125.184.0/22",
			"114.125.188.0/22",
			"114.125.192.0/22",
			"114.125.196.0/22",
			"114.125.200.0/22",
			"114.125.204.0/22",
			"114.125.208.0/22",
			"114.125.212.0/22",
			"114.125.216.0/22",
			"114.125.220.0/22",
			"114.125.224.0/22",
			"114.125.228.0/22",
			"114.125.232.0/22",
			"114.125.236.0/22",
			"114.125.240.0/22",
			"114.125.244.0/22",
			"114.125.248.0/22",
			"114.125.252.0/22",
			"114.126.0.0/22",
			"114.126.4.0/22",
			"114.126.8.0/22",
			"114.126.12.0/22",
			"114.126.16.0/22",
			"114.126.20.0/22",
			"114.126.24.0/22",
			"114.126.28.0/22",
			"114.126.32.0/22",
			"114.126.36.0/22",
			"114.126.40.0/22",
			"114.126.44.0/22",
			"114.126.48.0/22",
			"114.126.52.0/22",
			"114.126.56.0/22",
			"114.126.60.0/22",
			"114.126.64.0/22",
			"114.126.68.0/22",
			"114.126.72.0/22",
			"114.126.76.0/22",
			"114.126.80.0/22",
			"114.126.84.0/22",
			"114.126.88.0/22",
			"114.126.92.0/22",
			"114.126.96.0/22",
			"114.126.100.0/22",
			"114.126.104.0/22",
			"114.126.108.0/22",
			"114.126.112.0/22",
			"114.126.116.0/22",
			"114.126.120.0/22",
			"114.126.124.0/22",
			"114.126.128.0/22",
			"114.126.132.0/22",
			"114.126.136.0/22",
			"114.126.140.0/22",
			"114.126.144.0/22",
			"114.126.148.0/22",
			"114.126.152.0/22",
			"114.126.156.0/22",
			"114.126.160.0/22",
			"114.126.164.0/22",
			"114.126.168.0/22",
			"114.126.172.0/22",
			"114.126.176.0/22",
			"114.126.180.0/22",
			"114.126.184.0/22",
			"114.126.188.0/22",
			"114.126.192.0/22",
			"114.126.196.0/22",
			"114.126.200.0/22",
			"114.126.204.0/22",
			"114.126.208.0/22",
			"114.126.212.0/22",
			"114.126.216.0/22",
			"114.126.220.0/22",
			"114.126.224.0/22",
			"114.126.228.0/22",
			"114.126.232.0/22",
			"114.126.236.0/22",
			"114.126.240.0/22",
			"114.126.244.0/22",
			"114.126.248.0/22",
			"114.126.252.0/22",
			"114.127.0.0/22",
			"114.127.4.0/22",
			"114.127.8.0/22",
			"114.127.12.0/22",
			"114.127.16.0/22",
			"114.127.20.0/22",
			"114.127.24.0/22",
			"114.127.28.0/22",
			"114.127.32.0/22",
			"114.127.36.0/22",
			"114.127.40.0/22",
			"114.127.44.0/22",
			"114.127.48.0/22",
			"114.127.52.0/22",
			"114.127.56.0/22",
			"114.127.60.0/22",
			"114.127.64.0/22",
			"114.127.68.0/22",
			"114.127.72.0/22",
			"114.127.76.0/22",
			"114.127.80.0/22",
			"114.127.84.0/22",
			"114.127.88.0/22",
			"114.127.92.0/22",
			"114.127.96.0/22",
			"114.127.100.0/22",
			"114.127.104.0/22",
			"114.127.108.0/22",
			"114.127.112.0/22",
			"114.127.116.0/22",
			"114.127.120.0/22",
			"114.127.124.0/22",
			"114.127.128.0/22",
			"114.127.132.0/22",
			"114.127.136.0/22",
			"114.127.140.0/22",
			"114.127.144.0/22",
			"114.127.148.0/22",
			"114.127.152.0/22",
			"114.127.156.0/22",
			"114.127.160.0/22",
			"114.127.164.0/22",
			"114.127.168.0/22",
			"114.127.172.0/22",
			"114.127.176.0/22",
			"114.127.180.0/22",
			"114.127.184.0/22",
			"114.127.188.0/22",
			"114.127.192.0/22",
			"114.127.196.0/22",
			"114.127.200.0/22",
			"114.127.204.0/22",
			"114.127.208.0/22",
			"114.127.212.0/22",
			"114.127.216.0/22",
			"114.127.220.0/22",
			"114.127.224.0/22",
			"114.127.228.0/22",
			"114.127.232.0/22",
			"114.127.236.0/22",
			"114.127.240.0/22",
			"114.127.244.0/22",
			"114.127.248.0/22",
			"114.127.252.0/22",
			"114.122.0.0/19",
			"114.122.128.0/19",
			"114.122.160.0/19",
			"114.122.192.0/19",
			"114.122.32.0/19",
			"114.122.64.0/19",
		},
	},
	"my": {
		Hni:         "50212",
		CarrierName: "Maxis",
		CountryCode: "60",
	},
	"lt": {
		Hni:         "24601",
		CarrierName: "Omnitel",
		CountryCode: "370",
	},
	"tw": {
		Hni:         "46692",
		CarrierName: "Chunghwa",
		CountryCode: "886",
	},
	"us": {
		Hni:         "310012",
		CarrierName: "Verizon",
		CountryCode: "1",
		CIDRs: []string{
			"162.190.0.0/16",
			"162.191.0.0/16",
			"167.20.0.0/16",
			"172.56.0.0/16",
			"107.106.0.0/18",
			"107.106.120.0/22",
			"107.106.124.0/22",
			"107.106.128.0/18",
			"107.106.184.0/22",
			"107.106.188.0/22",
			"107.106.192.0/18",
			"107.106.56.0/22",
			"107.106.60.0/22",
			"107.106.64.0/18",
			"107.112.128.0/19",
			"107.112.160.0/19",
			"107.117.64.0/20",
			"107.117.80.0/20",
			"107.121.148.0/22",
			"107.121.156.0/22",
			"107.121.168.0/22",
			"107.124.192.0/18",
			"107.239.64.0/20",
			"107.239.80.0/20",
			"158.26.0.0/21",
		},
	},
	"hk": {
		Hni:         "45406",
		CarrierName: "SmarTone",
		CountryCode: "852",
	},
	"kg": {
		Hni:         "43705",
		CarrierName: "MegaCom",
		CountryCode: "996",
	},
	"no": {
		Hni:         "24204",
		CarrierName: "Tele2",
		CountryCode: "47",
	},
	"tr": {
		Hni:         "28603",
		CarrierName: "Avea",
		CountryCode: "90",
	},
	"es": {
		Hni:         "21401",
		CarrierName: "Vodafone",
		CountryCode: "34",
	},
	"gy": {
		Hni:         "73801",
		CarrierName: "Digicel",
		CountryCode: "592",
		CIDRs: []string{
			"181.177.216.0/22",
			"181.177.216.0/24",
			"181.177.217.0/24",
			"181.177.218.0/24",
			"181.177.219.0/24",
		},
	},
	"fi": {
		Hni:         "24405",
		CarrierName: "Elisa",
		CountryCode: "358",
	},
	"dk": {
		Hni:         "23802",
		CarrierName: "Telenor",
		CountryCode: "45",
	},
	"ru": {
		Hni:         "25020",
		CarrierName: "Tele2",
		CountryCode: "7",
		CIDRs: []string{
			"176.59.96.0/19",
			"93.93.136.0/21",
			"94.240.112.0/21",
			"94.240.64.0/18",
			"176.59.160.0/19",
			"176.59.160.0/22",
			"176.59.164.0/22",
			"176.59.168.0/22",
			"176.59.172.0/22",
			"176.59.32.0/19",
			"176.59.32.0/22",
			"176.59.36.0/22",
			"176.59.40.0/22",
		},
	},
	"th": {
		Hni:         "52015",
		CarrierName: "TOT 3G",
		CountryCode: "66",
		CIDRs: []string{
			"101.108.0.0/15",
			"101.108.0.0/17",
			"101.108.0.0/18",
			"101.108.0.0/19",
			"101.108.0.0/22",
			"101.108.112.0/21",
			"101.108.120.0/22",
			"101.108.12.0/22",
			"101.108.124.0/23",
			"101.108.126.0/23",
			"101.108.128.0/17",
			"101.108.128.0/18",
			"101.108.128.0/19",
			"101.108.128.0/21",
			"101.108.136.0/21",
			"101.108.144.0/21",
			"101.108.160.0/19",
			"101.108.160.0/20",
			"101.108.16.0/20",
			"101.108.176.0/20",
			"101.108.192.0/18",
			"101.108.192.0/19",
			"101.108.192.0/21",
			"101.108.200.0/21",
			"101.108.208.0/21",
			"101.108.216.0/22",
			"101.108.224.0/19",
			"101.108.224.0/20",
			"101.108.240.0/20",
			"101.108.32.0/19",
			"101.108.32.0/20",
			"101.108.4.0/22",
			"101.108.48.0/21",
			"101.108.56.0/22",
			"101.108.64.0/18",
			"101.108.64.0/19",
			"101.108.64.0/20",
			"101.108.80.0/22",
			"101.108.8.0/22",
			"101.108.96.0/19",
			"101.108.96.0/20",
		},
	},
	"in": {
		Hni:         "405025",
		CarrierName: "TATA Teleservice",
		CountryCode: "91",
	},
	"vn": {
		Hni:         "45201",
		CarrierName: "MobiFone",
		CountryCode: "84",
		CIDRs: []string{
			"103.19.96.0/22",
			"103.19.98.0/23",
			"103.23.156.0/22",
			"103.23.156.0/23",
			"103.23.158.0/23",
			"103.249.20.0/22",
			"103.249.20.0/23",
			"103.249.22.0/23",
			"103.7.36.0/22",
			"103.7.36.0/23",
			"103.7.38.0/23",
			"202.79.232.0/21",
			"202.79.232.0/23",
			"202.79.234.0/23",
			"202.79.236.0/23",
			"202.79.238.0/23",
			"203.170.26.0/23",
			"59.153.212.0/23",
			"59.153.214.0/23",
		},
	},
	"phi": {
		Hni:         "51502",
		CarrierName: "Globe",
		CountryCode: "63",
		CIDRs: []string{
			"110.54.128.0/24",
			"110.54.129.0/24",
			"110.54.130.0/24",
			"110.54.131.0/24",
			"110.54.136.0/24",
			"110.54.137.0/24",
			"110.54.138.0/24",
			"110.54.139.0/24",
			"110.54.144.0/24",
			"110.54.145.0/24",
			"110.54.146.0/24",
			"110.54.147.0/24",
			"110.54.160.0/24",
			"110.54.161.0/24",
			"110.54.162.0/24",
			"110.54.163.0/24",
			"110.54.164.0/24",
			"110.54.176.0/24",
			"110.54.177.0/24",
			"110.54.178.0/24",
			"110.54.179.0/24",
			"110.54.180.0/24",
			"110.54.192.0/24",
			"110.54.193.0/24",
			"110.54.194.0/24",
			"110.54.195.0/24",
			"110.54.196.0/24",
			"110.54.200.0/24",
			"110.54.201.0/24",
			"110.54.202.0/24",
			"110.54.203.0/24",
			"110.54.208.0/24",
			"110.54.209.0/24",
			"110.54.210.0/24",
			"110.54.211.0/24",
			"110.54.212.0/24",
			"110.54.215.0/24",
			"110.54.216.0/24",
			"110.54.217.0/24",
			"110.54.218.0/24",
			"110.54.219.0/24",
			"110.54.220.0/24",
			"110.54.224.0/24",
			"110.54.225.0/24",
			"110.54.226.0/24",
			"110.54.227.0/24",
			"110.54.232.0/24",
			"110.54.233.0/24",
			"110.54.234.0/24",
			"110.54.235.0/24",
			"110.54.236.0/24",
			"110.54.237.0/24",
			"110.54.238.0/24",
			"110.54.239.0/24",
			"110.54.240.0/24",
			"110.54.241.0/24",
			"110.54.242.0/24",
			"110.54.243.0/24",
			"110.54.244.0/24",
			"110.54.245.0/24",
			"110.54.246.0/24",
			"110.54.247.0/24",
			"110.54.248.0/24",
			"110.54.249.0/24",
			"110.54.250.0/24",
			"110.54.251.0/24",
			"111.90.200.0/24",
			"111.90.201.0/24",
			"111.90.202.0/24",
			"111.90.203.0/24",
			"111.90.204.0/24",
			"111.90.205.0/24",
			"111.90.206.0/24",
			"111.90.207.0/24",
			"111.90.208.0/24",
			"111.90.209.0/24",
			"111.90.210.0/24",
			"111.90.211.0/24",
			"111.90.212.0/24",
			"111.90.213.0/24",
		},
	},
	"jp": {
		Hni:         "44010",
		CarrierName: "NTT docomo",
		CountryCode: "81",
		CIDRs: []string{
			"101.102.0.0/18",
			"103.124.0.0/22",
			"110.165.128.0/17",
			"119.30.192.0/18",
			"119.31.128.0/19",
			"133.106.128.0/17",
			"133.106.16.0/20",
			"133.106.32.0/19",
			"133.106.64.0/18",
			"133.106.8.0/21",
			"157.192.0.0/16",
			"193.114.192.0/18",
			"193.114.32.0/19",
			"193.114.64.0/19",
			"193.115.0.0/19",
			"193.117.96.0/19",
			"193.118.0.0/19",
			"193.118.64.0/19",
			"193.119.128.0/17",
			"193.82.160.0/19",
			"194.193.224.0/19",
			"194.193.64.0/19",
			"194.223.96.0/19",
			"202.176.16.0/20",
			"202.216.0.0/20",
			"210.157.192.0/19",
			"211.133.160.0/19",
			"211.7.96.0/19",
			"219.105.144.0/20",
			"219.105.192.0/18",
			"219.106.0.0/17",
		},
	},
}
