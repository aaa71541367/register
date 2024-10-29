package helper

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"reflect"
	"strconv"
	"time"
	"unsafe"

	"github.com/shillbie/register/LineThrift"
	"github.com/valyala/fasthttp"

	"github.com/utahta/go-linenotify"
)

type mentionMsg struct {
	MENTIONEES []struct {
		S string `json:"S"`
		E string `json:"E"`
		M string `json:"M"`
	} `json:"MENTIONEES"`
}

func FmtDuration(d time.Duration) string {
	d = d.Round(time.Second)
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	d -= m * time.Minute
	s := d / time.Second
	return fmt.Sprintf("%02d days %02d hours %02d min %02d sec.", h/24, h%24, m, s)
}

func GetMentionMsg() mentionMsg {
	return mentionMsg{}
}

func MaxRevision(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func ArrayCast(src, dst interface{}, cap, len int) {
	sh := (*reflect.SliceHeader)(unsafe.Pointer(reflect.ValueOf(dst).Pointer()))
	sh.Cap, sh.Len = cap, len
	sh.Data = uintptr(unsafe.Pointer(reflect.ValueOf(src).Pointer()))
}

func SetField(obj interface{}, name string, value interface{}) error {
	structValue := reflect.ValueOf(obj).Elem()
	structFieldValue := structValue.FieldByName(name)

	if !structFieldValue.IsValid() {
		return fmt.Errorf("No such field: %s in obj", name)
	}

	if !structFieldValue.CanSet() {
		return fmt.Errorf("Cannot set %s field value", name)
	}

	structFieldType := structFieldValue.Type()
	val := reflect.ValueOf(value)
	if structFieldType != val.Type() {
		return errors.New("Provided value type didn't match obj field type")
	}

	structFieldValue.Set(val)
	return nil
}

func InArray(array interface{}, val interface{}) bool {
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(array)
		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				return true
			}
		}
	}
	return false
}

func Append(arr []string, strs ...string) []string {
OUTER:
	for _, str := range strs {
		for _, a := range arr {
			if a == str {
				continue OUTER
			}
		}
		arr = append(arr, str)
	}
	return arr
}

func InMap(dict map[string]bool, key string) bool {
	_, ok := dict[key]
	return ok
}

func Zip(a []interface{}, b []string) ([]interface{}, error) {
	if len(a) != len(b) {
		return nil, fmt.Errorf("zip: arguments must be of same length")
	}
	r := make([]interface{}, len(a))
	for i, e := range a {
		r[i] = []interface{}{e, b[i]}
	}
	return r, nil
}

func IndexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1
}

func Remove(s []string, r string) []string {
	new := make([]string, len(s))
	copy(new, s)
	for i, v := range new {
		if v == r {
			return append(new[:i], new[i+1:]...)
		}
	}
	return s
}

func SplitList(s []string, a int) [][]string {
	t, r := len(s)/a, len(s)%a
	res := [][]string{}
	for x := 0; x < t; x++ {
		res = append(res, s[a*x:a*(x+1)])
	}
	if r != 0 {
		res = append(res, s[a*t:])
	}
	return res
}

func GroupList(s []string, a int) [][]string {
	res := make([][]string, a)
	for no, i := range s {
		res[no%a] = append(res[no%a], i)
	}
	return res
}

func RandChoice(data []string) string {
	rand.Seed(time.Now().UTC().UnixNano())
	return data[rand.Intn(len(data))]
}

func MatchList(a []string, b []string) []string {
	res := []string{}
	for _, v := range a {
		if InArray(b, v) && !InArray(res, v) {
			res = append(res, v)
		}
	}
	return res
}

func Shift(s []string, a int) []string {
	if a >= len(s) || a < 0 {
		fmt.Println("Invalid number of length.")
		return s
	}
	x := s[:a]
	y := s[a:]
	return append(y, x...)
}

func GetMidFromMentionees(data string) []string {
	var midmen []string
	var midbefore []string
	res := mentionMsg{}
	json.Unmarshal([]byte(data), &res)
	for _, v := range res.MENTIONEES {
		if InArray(midbefore, v.M) == false {
			midbefore = append(midbefore, v.M)
			midmen = append(midmen, v.M)
		}
	}
	return midmen
}

func Log(optype LineThrift.OpType, logtype string, str string) {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	a := time.Now().In(loc)
	yyyy := strconv.Itoa(a.Year())
	MM := a.Month().String()
	dd := strconv.Itoa(a.Day())
	hh := a.Hour()
	mm := a.Minute()
	ss := a.Second()
	var hhconv string
	var mmconv string
	var ssconv string
	if hh < 10 {
		hhconv = "0" + strconv.Itoa(hh)
	} else {
		hhconv = strconv.Itoa(hh)
	}
	if mm < 10 {
		mmconv = "0" + strconv.Itoa(mm)
	} else {
		mmconv = strconv.Itoa(mm)
	}
	if ss < 10 {
		ssconv = "0" + strconv.Itoa(ss)
	} else {
		ssconv = strconv.Itoa(ss)
	}
	times := yyyy + "-" + MM + "-" + dd + " " + hhconv + ":" + mmconv + ":" + ssconv
	fmt.Println("[" + times + "][" + optype.String() + "][" + logtype + "]" + str)
}

func NotifySendMessage(text string) error {
	var notify = linenotify.NewClient()
	_, err := notify.Notify(context.TODO(), "ldx4Xt442DBv4Xc9HQIHqoSkCURlkR7zinJoKRHmy8R", "\n"+text, "", "", nil)
	return err
}

func DownloadFile(url string) (path string, err error) {
	path, err = os.Getwd()
	if err != nil {
		return path, err
	}
	tmpfile, err := ioutil.TempFile(path+"/", "LineAR-GOBOT-")
	if err != nil {
		return path, err
	}
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	req.SetRequestURI(url)
	req.Header.SetMethod("GET")
	err = fasthttp.DoRedirects(req, resp, 10)
	if err != nil {
		return path, err
	}
	tmpfile.Write(resp.Body())
	if err != nil {
		return path, err
	}
	defer tmpfile.Close()
	return tmpfile.Name(), nil
}
