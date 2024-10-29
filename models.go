package goline

import (
	"bytes"
	"crypto/md5"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path/filepath"

	//"strconv"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

/* files */

func (p *LINE) SaveFile(path string, data io.Reader) error {
	f, err := os.Create("data.txt")
	if err != nil {
		return err
	}
	_, err = io.Copy(f, data)
	if err != nil {
		return err
	}
	defer f.Close()
	return nil
}

func (p *LINE) DeleteFile(path string) error {
	return os.Remove(path)
}

func (p *LINE) DownloadFile(url string) (path string, err error) {
	path, err = os.Getwd()
	if err != nil {
		return path, err
	}
	tmpfile, err := ioutil.TempFile(path+"/", "LineAR-GOBOT-")
	if err != nil {
		return path, err
	}
	res, err := p.GetContent(url, []byte{}, nil)
	if err != nil {
		return path, err
	}
	if _, err := tmpfile.Write(res.Body()); err != nil {
		return path, err
	}
	return tmpfile.Name(), nil
}

func (p *LINE) DownloadFromMsg(msgid string) (string, error) {
	return p.DownloadFile(LINE_OBS_DOMAIN + "/talk/m/download.nhn?oid=" + msgid)
}

func (p *LINE) UpdateProfilePicture(name string) (err error) {
	if p.timeLineHeader == nil || p.timeLineHeader["X-Line-ChannelToken"] == "" {
		p.SetTimelineHeaders()
	}
	hstr := fmt.Sprintf("Amemiya_%d000", time.Now().Unix())
	objid := fmt.Sprintf("%x", md5.Sum([]byte(hstr)))
	url := LINE_OBS_DOMAIN + "/r/talk/p/" + p.MID

	file, err := os.Open(name)
	if err != nil {
		return err
	}
	defer file.Close()
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormField("params")
	params, _ := json.Marshal(map[string]string{
		"name":    objid,
		"quality": "100",
		"type":    "image",
		"ver":     "2.0",
	})
	part.Write(params)
	part, err = writer.CreateFormFile("file", filepath.Base(name))
	if err != nil {
		return err
	}
	_, err = io.Copy(part, file)
	err = writer.Close()
	if err != nil {
		return err
	}
	extraHeader := map[string]string{
		"Content-Type":        writer.FormDataContentType(),
		"X-Line-Mid":          p.MID,
		"X-Line-ChannelToken": p.timeLineHeader["X-Line-ChannelToken"],
	}
	_, err = p.PostContent(url, body.Bytes(), extraHeader)
	return err
}

func (p *LINE) UpdateProfilePicture2(name string) (err error) {
	url := LINE_OBS_DOMAIN + "/r/talk/p/" + p.MID
	file, err := os.Open(name)
	if err != nil {
		return err
	}
	defer file.Close()
	bData, _ := ioutil.ReadAll(file)
	params, _ := json.Marshal(map[string]string{
		"name": "ImageProfile",
		"oid":  p.MID,
		"type": "image",
		"ver":  "1.0",
	})
	encToken, _ := p.TalkService().AcquireEncryptedAccessToken(p.ctx, 2)
	extraHeader := map[string]string{
		"x-lal":          "zh-Hant_TW",
		"X-Line-Access":  encToken,
		"x-line-carrier": "46689, 1",
		"Content-Type":   "application/octet-stream",
		"x-obs-Params":   base64.StdEncoding.EncodeToString(params),
		"Cookie":         "os=ios",
	}
	_, err = p.PostContent(url, bData, extraHeader)
	return err
}

func (p *LINE) UpdateGroupPicture(to string, name string) (err error) {
	file, err := os.Open(name)
	if err != nil {
		return err
	}
	defer file.Close()
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormField("params")
	params, _ := json.Marshal(map[string]string{
		"name": "line-ar.bin",
		"ver":  "1.0",
		"oid":  to,
		"type": "image",
	})
	part.Write(params)
	part, err = writer.CreateFormFile("file", filepath.Base(name))
	if err != nil {
		return err
	}
	_, err = io.Copy(part, file)
	err = writer.Close()
	if err != nil {
		return err
	}
	extraHeader := map[string]string{
		"Content-Type": writer.FormDataContentType(),
	}
	_, err = p.PostContent(LINE_OBS_DOMAIN+"/talk/g/upload.nhn", body.Bytes(), extraHeader)
	p.DeleteFile(name)
	return err
}

func (p *LINE) UploadObjTalk(name string, contentType string, objid string) (err error) {
	if contentType == "image" || contentType == "video" || contentType == "audio" || contentType == "file" {
		file, err := os.Open(name)
		if err != nil {
			return err
		}
		defer file.Close()
		fileInfo, _ := file.Stat()
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		part, err := writer.CreateFormField("params")
		params, _ := json.Marshal(map[string]interface{}{
			"name": "line-ar.bin",
			"ver":  "1.0",
			"oid":  objid,
			"type": contentType,
			"size": fileInfo.Size(),
		})
		part.Write(params)
		part, err = writer.CreateFormFile("file", filepath.Base(name))
		if err != nil {
			return err
		}
		_, err = io.Copy(part, file)
		err = writer.Close()
		if err != nil {
			return err
		}
		extraHeader := map[string]string{
			"Content-Type":        writer.FormDataContentType(),
			"X-Line-Mid":          p.MID,
			"X-Line-ChannelToken": p.timeLineHeader["X-Line-ChannelToken"],
		}
		_, err = p.PostContent(LINE_OBS_DOMAIN+"/talk/m/upload.nhn", body.Bytes(), extraHeader)

		return err
	} else {
		return fmt.Errorf("Invalid type")
	}
}

func (p *LINE) UploadObjSquare(squareChatMid, path, contentType string) error {
	p.reqSeq += 1
	if contentType == "image" || contentType == "video" || contentType == "audio" {
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()
		fileInfo, _ := file.Stat()
		bData, err := ioutil.ReadAll(file)
		param, _ := json.Marshal(map[string]interface{}{
			"ver":    "2.0",
			"name":   fileInfo.Name(),
			"oid":    "reqseq",
			"tomid":  squareChatMid,
			"reqseq": fmt.Sprintf("%d", p.reqSeq),
			"type":   contentType,
		})
		extraHeader := map[string]string{
			"Content-Type": "application/octet-stream",
			"x-obs-params": base64.StdEncoding.EncodeToString(param),
		}
		_, err = p.PostContent(LINE_OBS_DOMAIN+"/r/g2/m/reqseq", bData, extraHeader)
		return err
	}
	return fmt.Errorf("Invalid contentType")
}

func (p *LINE) UploadObjsSquare(squareChatMid string, paths []string, contentType string) error {
	createMeta := func(data map[string]string) string {
		sqrd := []byte{10, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 10, 0, 6, 0, 0, 0, 0, 0, 0, 0, 0, 2, 0, 14, 0, 13, 0, 18, 11, 11}
		sqrd = append(sqrd, []byte{0, 0, 0, 3}...)
		sqrd = append(sqrd, []byte{0, 0, 0, 3, 71, 73, 68, 0, 0, 0, byte(len(data["GID"]))}...)
		sqrd = append(sqrd, []byte(data["GID"])...)
		sqrd = append(sqrd, []byte{0, 0, 0, 6, 71, 84, 79, 84, 65, 76, 0, 0, 0, byte(len(data["GTOTAL"]))}...)
		sqrd = append(sqrd, []byte(data["GTOTAL"])...)
		sqrd = append(sqrd, []byte{0, 0, 0, 4, 71, 83, 69, 81, 0, 0, 0, byte(len(data["GSEQ"]))}...)
		sqrd = append(sqrd, []byte(data["GSEQ"])...)
		sqrd = append(sqrd, []byte{3, 0, 19, 0, 15, 0, 27, 12, 0, 0, 0, 0, 0}...)
		bData, _ := json.Marshal(map[string]string{"message": base64.StdEncoding.EncodeToString(sqrd)})
		return base64.StdEncoding.EncodeToString(bData)
	}
	if contentType == "image" || contentType == "video" || contentType == "audio" {
		gid := "0"
		for no, path := range paths {
			p.reqSeq += 1
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			// fileInfo, _ := file.Stat()
			bData, err := ioutil.ReadAll(file)
			param, _ := json.Marshal(map[string]interface{}{
				"ver":     "1.0",
				"name":    fmt.Sprintf("%d", p.reqSeq),
				"oid":     "reqseq",
				"tomid":   squareChatMid,
				"reqseq":  fmt.Sprintf("%d", p.reqSeq),
				"type":    contentType,
				"quality": "70",
			})
			extraHeader := map[string]string{
				"X-Square-Meta": createMeta(map[string]string{"GID": gid, "GSEQ": fmt.Sprintf("%d", no+1), "GTOTAL": fmt.Sprintf("%d", len(paths))}),
				"Content-Type":  "application/octet-stream",
				"x-obs-params":  base64.StdEncoding.EncodeToString(param),
			}
			if res, err := p.PostContent(LINE_OBS_DOMAIN+"/r/g2/m/reqseq", bData, extraHeader); err != nil {
				return err
			} else {
				gid = string(res.Header.Peek("x-line-message-gid"))
			}
		}
		return nil
	}
	return fmt.Errorf("Invalid contentType")
}

func (p *LINE) UploadObjHome(path string, contentType string) (objId string, err error) {
	if contentType != "image/jpeg" && contentType != "video/mp4" && contentType != "audio/mp3" {
		return "", fmt.Errorf("Invalid type")
	} else if p.timeLineHeader == nil || p.timeLineHeader["X-Line-ChannelToken"] == "" {
		p.SetTimelineHeaders()
	}
	objId = fmt.Sprintf("%d", time.Now().UnixNano()/int64(time.Second))
	file, err := os.Open(path)
	if err != nil {
		return objId, err
	}
	defer file.Close()
	fileInfo, _ := file.Stat()
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filepath.Base(path))
	if err != nil {
		return objId, err
	}
	_, err = io.Copy(part, file)
	err = writer.Close()
	param, _ := json.Marshal(map[string]string{
		"name":   "line-ar.bin",
		"ver":    "1.0",
		"userid": p.MID,
		"oid":    objId,
		"range":  fmt.Sprintf("%d", fileInfo.Size()),
		"type":   contentType,
	})
	extraHeader := map[string]string{
		"Content-Type":        writer.FormDataContentType(),
		"Content-Length":      fmt.Sprintf("%d", fileInfo.Size()),
		"X-Line-Mid":          p.MID,
		"X-Line-ChannelToken": p.timeLineHeader["X-Line-ChannelToken"],
		"x-obs-params":        base64.StdEncoding.EncodeToString(param),
	}
	_, err = p.PostContent(LINE_OBS_DOMAIN+"/myhome/c/upload.nhn", body.Bytes(), extraHeader)
	return objId, err
}

func (p *LINE) UpdateProfileCover(path string) (err error) {
	if p.timeLineHeader == nil || p.timeLineHeader["X-Line-ChannelToken"] == "" {
		p.SetTimelineHeaders()
	}
	hstr := fmt.Sprintf("Amemiya_%d000", time.Now().Unix())
	objid := fmt.Sprintf("%x", md5.Sum([]byte(hstr)))
	url := LINE_OBS_DOMAIN + "/r/myhome/c/" + objid

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormField("params")
	params, _ := json.Marshal(map[string]string{
		"name":    objid,
		"quality": "100",
		"type":    "image",
		"ver":     "2.0",
	})
	part.Write(params)
	part, err = writer.CreateFormFile("file", filepath.Base(path))
	if err != nil {
		return err
	}
	_, err = io.Copy(part, file)
	err = writer.Close()
	if err != nil {
		return err
	}
	extraHeader := map[string]string{
		"Content-Type":        writer.FormDataContentType(),
		"X-Line-Mid":          p.MID,
		"X-Line-ChannelToken": p.timeLineHeader["X-Line-ChannelToken"],
	}
	_, err = p.PostContent(url, body.Bytes(), extraHeader)
	p.UpdateProfileCoverById(objid)
	return err
}

func (p *LINE) UpdateProfileCover2(name string) (err error) {
	if p.timeLineHeader == nil || p.timeLineHeader["X-Line-ChannelToken"] == "" {
		p.SetTimelineHeaders()
	}
	hstr := fmt.Sprintf("Amemiya_%d", time.Now().Unix())
	objid := fmt.Sprintf("%x", md5.Sum([]byte(hstr)))
	url := LINE_OBS_DOMAIN + "/r/myhome/c/" + objid
	file, err := os.Open(name)
	if err != nil {
		return err
	}
	defer file.Close()
	bData, _ := ioutil.ReadAll(file)
	params, _ := json.Marshal(map[string]string{
		"name":   objid + ".jpg",
		"userid": p.MID,
		"oid":    objid,
		"type":   "image",
		"ver":    "1.0",
	})
	extraHeader := map[string]string{
		"x-lal":          "zh-Hant_TW",
		"x-line-carrier": "46689, 1",
		"Content-Type":   "application/octet-stream",
		"x-obs-params":   base64.StdEncoding.EncodeToString(params),
		"Cookie":         "os=ios",
	}
	_, err = p.TimelineRequests(url, "POST", extraHeader, bData)
	err = p.UpdateProfileCoverById2(objid)
	return err
}

func (p *LINE) UpdateProfileCoverById(objid string) (err error) {
	data, _ := json.Marshal(map[string]interface{}{
		"homeId":        p.MID,
		"coverObjectId": objid,
		"storyShare":    true,
		"meta":          map[string]string{},
	})
	req, err := http.NewRequest("POST", LINE_HOST_DOMAIN+"/hm/api/v1/home/cover.json", bytes.NewBuffer(data))
	for k, v := range p.timeLineHeader {
		req.Header.Set(k, v)
	}
	req.Header.Set("x-lhm", "POST")
	_, err = new(http.Client).Do(req)
	return err
}

func (p *LINE) UpdateProfileCoverById2(objid string) (err error) {
	extraHeader := map[string]string{
		"x-lal":                "zh-Hant_TW",
		"x-lhm":                "POST",
		"Content-Type":         "application/json",
		"x-line-signup-region": "TW",
		"x-line-global-config": "reboot.phase=scenario; discover.enable=true; follow.enable=true",
		"x-lpv":                "1",
	}
	_, err = p.TimelineRequests(LINE_HOST_DOMAIN+"/hm/api/v1/home/cover.json", "POST", extraHeader, map[string]interface{}{
		"homeId": p.MID,
		"storyMeta": map[string]string{
			"storyVersion": "v8",
		},
		"storyShare":    false,
		"coverObjectId": objid,
	})
	return err
}
