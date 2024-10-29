package goline

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	"github.com/tidwall/gjson"
	"github.com/valyala/fasthttp"
)

func (p *LINE) SetTimelineHeaders() error {
	channel, err := p.ChannelService().ApproveChannelAndIssueChannelToken(p.ctx, TIMELINE_CHANNEL_ID)
	if err != nil {
		return err
	}
	p.timeLineHeader = map[string]string{
		"Content-Type":        "application/json",
		"X-LAP":               "5",
		"X-LPV":               "1",
		"X-Line-ChannelToken": channel.ChannelAccessToken,
		"X-Line-Application":  p.AppName,
		"User-Agent":          p.userAgent,
		"X-Line-Mid":          p.MID,
	}
	// fmt.Println("ChannelToken: ", channel.ChannelAccessToken)
	return nil
}

func (p *LINE) TimelineRequests(url string, method string, extraHeader map[string]string, data interface{}) (res string, err error) {
	if p.timeLineHeader == nil || p.timeLineHeader["X-Line-ChannelToken"] == "" {
		if err = p.SetTimelineHeaders(); err != nil {
			return "", err
		}
	}
	if p.fast_connection == nil {
		p.setDefaultHttpClient()
	}
	request := fasthttp.AcquireRequest()
	response := fasthttp.AcquireResponse()
	request.SetRequestURI(url)
	request.Header.SetMethod(method)
	for key, value := range p.timeLineHeader {
		request.Header.Set(key, value)
	}
	for key, value := range extraHeader {
		request.Header.Set(key, value)
	}
	if method == "POST" {
		if bData, ok := data.([]byte); ok {
			request.SetBody(bData)
		} else {
			byteData, _ := json.Marshal(data)
			request.SetBody(byteData)
		}
	}
	err = p.fast_connection.Do(request, response)
	if err != nil {
		return "", err
	}
	return string(response.Body()), nil
}

func (p *LINE) GetSocialPorfile(mid string) (string, error) {
	params := url.Values{}
	params.Add("homeId", mid)
	params.Add("storyVersion", "v8")
	params.Add("withSocialHomeInfo", "1")
	params.Add("timelineVersion", "v57")
	url := LINE_HOST_DOMAIN + "/hm/api/v2/home/socialprofile/post?" + params.Encode()
	r, err := p.TimelineRequests(url, "GET", nil, "")
	if gjson.Get(r, "code").Int() != 0 {
		return "", fmt.Errorf(gjson.Get(r, "message").String())
	}
	return r, err
}

func (p *LINE) GetProfileDetail(mid string) {
	params := url.Values{}
	params.Add("homeId", mid)
	params.Add("styleMediaVersion", "v2")
	params.Add("storyVersion", "v7")
	params.Add("timelineVersion", "v57")
	params.Add("profileBannerRevision", "0")
	url := LINE_HOST_DOMAIN + "/hm/api/v1/home/profile.json?" + params.Encode()
	r, err := p.TimelineRequests(url, "GET", nil, "")
	fmt.Println(r, err)
}

func (p *LINE) GetProfileCoverDetail(mid string) (string, error) {
	r, err := p.TimelineRequests(LINE_HOST_DOMAIN+"/hm/api/v1/home/cover.json?homeId="+mid, "GET", nil, "")
	return r, err
}

func (p *LINE) UpdateCover(objId string) (string, error) {
	url := LINE_HOST_DOMAIN + "/mh/api/v39/home/updateCover.json?coverImageId=" + objId
	r, err := p.TimelineRequests(url, "GET", nil, "")
	return r, err
}

func (p *LINE) DeletePost(pid string, to string) (string, error) {
	if to == "" {
		to = p.MID
	}
	params := url.Values{}
	params.Add("homeId", to)
	params.Add("postId", pid)
	url := LINE_HOST_DOMAIN + "/mh/v48/post/delete.json?" + params.Encode()
	r, err := p.TimelineRequests(url, "GET", nil, "")
	return r, err
}

func (p *LINE) SendPostToTalk(mid string, postId string) (string, error) {
	params := url.Values{}
	params.Add("receiveMid", mid)
	params.Add("postId", postId)
	url := LINE_HOST_DOMAIN + "/mh/api/v39/post/sendPostToTalk.json?" + params.Encode()
	r, err := p.TimelineRequests(url, "GET", nil, "")
	return r, err
}

func (p *LINE) GetGroupPost(groupid string) (string, error) {
	params := url.Values{}
	params.Add("homeId", groupid)
	params.Add("postLimit", "100")
	params.Add("commentLimit", "50")
	params.Add("likeLimit", "500")
	params.Add("sourceType", "TALKROOM")
	url := LINE_HOST_DOMAIN + "/mh/api/v39/post/list.json?" + params.Encode()
	r, err := p.TimelineRequests(url, "GET", nil, "")
	return r, err
}

func (p *LINE) LikePost(postId string, likeType string) (string, error) {
	params := url.Values{}
	params.Add("homeId", p.MID)
	params.Add("sourceType", "TIMELINE")
	url := LINE_HOST_DOMAIN + "/mh/api/v39/like/create.json?" + params.Encode()
	datas := map[string]string{"likeType": likeType, "activityExternalId": postId, "actorId": p.MID}
	r, err := p.TimelineRequests(url, "POST", nil, datas)
	return r, err
}

func (p *LINE) CommantPost(postId string, text string) (string, error) {
	params := url.Values{}
	params.Add("homeId", p.MID)
	params.Add("sourceType", "TIMELINE")
	url := LINE_HOST_DOMAIN + "/mh/api/v39/comment/create.json?" + params.Encode()
	datas := map[string]string{"commentText": text, "activityExternalId": postId, "actorId": p.MID}
	r, err := p.TimelineRequests(url, "POST", nil, datas)
	return r, err
}

func (p *LINE) CreatePost(text string, mid string) (string, error) {
	params := url.Values{}
	params.Add("homeId", mid)
	params.Add("sourceType", "TIMELINE")
	url := LINE_HOST_DOMAIN + "/mh/api/v39/post/create.json?" + params.Encode()
	var data interface{}
	strdata := "{\"postInfo\": {\"readPermission\": {\"type\": \"ALL\"}}, \"sourceType\": \"TIMELINE\", \"contents\": {\"text\": \"" + text + "\"}}"
	json.Unmarshal([]byte(strdata), &data)
	r, err := p.TimelineRequests(url, "POST", nil, data)
	return r, err
}

func (p *LINE) CreateKeep(keepID, content string) (string, error) {
	if keepID == "" {
		keepID = fmt.Sprintf("%d", time.Now().UnixNano())
	}
	var data interface{}
	strdata := fmt.Sprintf(`{"clientId": "ad_kp|%s", "contentData": [{"size": 1, "text": "%s", "type": "TEXT", "urlList": [], "bgColor": "COLOR_04"}], "source": {"id": "ub7a12eba4acc71bd47e749ea46d8a09b", "type": "KEEP"}, "tagInfos": [], "collectionInfos": [], "pinInfo": {"pinned": false}, "serviceType": 0}`, keepID, content)
	json.Unmarshal([]byte(strdata), &data)
	r, err := p.TimelineRequests(LINE_HOST_DOMAIN+"/kp/api/v27/keep/create.json", "POST", nil, data)
	if gjson.Get(r, "code").Int() != 0 {
		return "", fmt.Errorf(gjson.Get(r, "message").String())
	}
	return r, err
}

func (p *LINE) DeleteKeeps(contentIds []string) (string, error) {
	data := struct {
		ContentIds []string `json:"contentIds"`
	}{
		ContentIds: contentIds,
	}
	r, err := p.TimelineRequests(LINE_HOST_DOMAIN+"/kp/api/v27/keep/delete.json", "POST", nil, data)
	return r, err
}

func (p *LINE) GetKeeps(revision, limit string) (string, error) {
	params := url.Values{}
	params.Add("revision", revision)
	params.Add("limit", limit)
	url := LINE_HOST_DOMAIN + "/kp/api/v27/keep/sync.json?" + params.Encode()
	r, err := p.TimelineRequests(url, "GET", nil, "")
	if gjson.Get(r, "code").Int() != 0 {
		return "", fmt.Errorf(gjson.Get(r, "message").String())
	}
	return r, err
}

func (p *LINE) ForwardObjectMsg(to, msgID, contentType string) (string, error) {
	url := LINE_OBS_DOMAIN + "/talk/m/copy.nhn"
	data := map[string]string{
		"name":     fmt.Sprintf("LINEGO-%d", time.Now().Unix()),
		"tomid":    to,
		"oid":      "reqseq",
		"reqseq":   fmt.Sprintf("%d", p.Revision),
		"type":     contentType,
		"copyFrom": "/talk/m/" + msgID,
	}
	r, err := p.TimelineRequests(url, "POST", nil, data)
	return r, err
}
