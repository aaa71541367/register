package helper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func GetTempName() (name string, err error) {
	res, err := http.Get("https://apis.tianapi.com/cname/index?key=3d9a5c4491c18e487526e3d09dc16aad&sex=2")
	if err != nil {
		return "", err
	}
	resp := struct {
		Code    int    `json:"code"`
		Message string `json:"msg"`
		Result  struct {
			List []struct {
				Name    string `json:"naming"`
				WordNum int    `json:"wordnum"`
				Sex     int    `json:"sex"`
			} `json:"list"`
		} `json:"result"`
	}{}
	err = json.NewDecoder(res.Body).Decode(&resp)
	if err != nil {
		return "", err
	}
	if resp.Code != 200 {
		return "", fmt.Errorf(resp.Message)
	} else {
		return resp.Result.List[0].Name, nil
	}
}

func GetTempStatusMessage() (string, error) {
	res, err := http.Get("https://apis.tianapi.com/hsjz/index?key=3d9a5c4491c18e487526e3d09dc16aad&sex=2")
	if err != nil {
		return "", err
	}
	resp := struct {
		Code    int    `json:"code"`
		Message string `json:"msg"`
		Result  struct {
			Content string `json:"content"`
		} `json:"result"`
	}{}
	err = json.NewDecoder(res.Body).Decode(&resp)
	if err != nil {
		return "", err
	}
	if resp.Code != 200 {
		return "", fmt.Errorf(resp.Message)
	} else {
		return resp.Result.Content, nil
	}
}

func GetTodayLock(astro string) (string, error) {
	res, err := http.Get("https://apis.tianapi.com/star/index?key=3d9a5c4491c18e487526e3d09dc16aad&astro=" + astro)
	if err != nil {
		return "", err
	}
	resp := struct {
		Code    int    `json:"code"`
		Message string `json:"msg"`
		Result  struct {
			List []struct {
				Type    string `json:"type"`
				Content string `json:"content"`
			} `json:"list"`
		} `json:"result"`
	}{}
	err = json.NewDecoder(res.Body).Decode(&resp)
	if err != nil {
		return "", err
	}
	if resp.Code != 200 {
		return "", fmt.Errorf(resp.Message)
	} else {
		res := "今日運勢"
		for _, ret := range resp.Result.List {
			res += "\n" + ret.Type + ":" + ret.Content
		}
		return res, nil
	}
}

func Simplify2Tradition(text string) string {
	form := url.Values{}
	form.Add("text", text)
	form.Add("config", "s2t.json")
	form.Add("precise", "0")
	req, _ := http.NewRequest("POST", "https://opencc.byvoid.com/convert", strings.NewReader(form.Encode()))
	req.Header.Set("User-Agent", "PostmanRuntime/7.29.2")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return text
	}
	defer resp.Body.Close()
	bData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return text
	}
	return string(bData)
}

func WhetherReport(location string) (string, error) {
	if InArray([]string{"宜蘭縣", "花蓮縣", "臺東縣", "澎湖縣", "金門縣", "連江縣", "臺北市", "新北市", "桃園市", "臺中市", "臺南市", "高雄市", "基隆市", "新竹縣", "新竹市", "苗栗縣", "彰化縣", "南投縣", "雲林縣", "嘉義縣", "嘉義市", "屏東縣"}, location) {
		res, err := http.Get("https://opendata.cwb.gov.tw/api/v1/rest/datastore/F-C0032-001?Authorization=CWB-AA786A07-E2FD-4ECE-8AD0-4677C6B293D6&format=JSON&locationName=" + location)
		if err != nil {
			return "", err
		}
		defer res.Body.Close()
		bData, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return "", err
		}
		resp := struct {
			Success string `json:"success"`
			Result  struct {
				ResourceID string `json:"resource_id"`
				Fields     []struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"fields"`
			} `json:"result"`
			Records struct {
				DatasetDescription string `json:"datasetDescription"`
				Location           []struct {
					LocationName   string `json:"locationName"`
					WeatherElement []struct {
						ElementName string `json:"elementName"`
						Time        []struct {
							StartTime string `json:"startTime"`
							EndTime   string `json:"endTime"`
							Parameter struct {
								ParameterName  string `json:"parameterName"`
								ParameterValue string `json:"parameterValue"`
							} `json:"parameter"`
						} `json:"time"`
					} `json:"weatherElement"`
				} `json:"location"`
			} `json:"records"`
		}{}
		err = json.Unmarshal(bData, &resp)
		if err != nil {
			return "", err
		}
		if resp.Success != "true" || len(resp.Records.Location) == 0 {
			return "", fmt.Errorf(string(bData))
		}
		ret := location + "未來36小時天氣預報"
		for _, elem := range resp.Records.Location[0].WeatherElement {
			switch elem.ElementName {
			case "PoP":
				ret += "\n下雨機率"
				for _, tData := range elem.Time {
					ret += "\n" + tData.StartTime + ":" + tData.Parameter.ParameterName + "%"
				}
			case "MinT":
				ret += "\n最低溫"
				for _, tData := range elem.Time {
					ret += "\n" + tData.StartTime + ":" + tData.Parameter.ParameterName + "°C"
				}
			}
		}
		return ret, nil
	} else {
		return "", fmt.Errorf("error location")
	}
}

func Covid19Data(country string) (string, error) {
	res, err := http.Get("https://covid-19.nchc.org.tw/api/covid19?CK=covid-19@nchc.org.tw&querydata=4003")
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	bData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	resp := []struct {
		ID  string `json:"id"`
		A01 string `json:"a01"`
		A02 string `json:"a02"`
		A03 string `json:"a03"`
		A04 string `json:"a04"`
		A05 string `json:"a05"`
		A06 string `json:"a06"`
		A07 string `json:"a07"`
		A08 string `json:"a08"`
		A09 string `json:"a09"`
	}{}
	err = json.Unmarshal(bData, &resp)
	if err != nil {
		return "", err
	}
	if len(resp) == 0 {
		return "", fmt.Errorf(string(bData))
	}
	ret := "今日Covid-19確診統計"
	ret += "\n送驗(今日合計)" + resp[0].A05
	ret += "\n送驗(總累計)" + resp[0].A06
	ret += "\n排除(總累計)" + resp[0].A07
	ret += "\n昨日排除" + resp[0].A08
	ret += "\n昨日送驗" + resp[0].A09
	return ret, nil
}
