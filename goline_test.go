package goline

import (
	"testing"
)

func TestLimit(t *testing.T) {
	me := NewLogin()
	me.Login("u07857d52aa4861d54cdd2ff85672e8ce:aWF0OiAxNTg2MTU1NzQ3OTczCg==..plZW8K62bG5RY/tfVnmLKbixNfA=")
	mids, err := me.GetAllContactIds()
	t.Log(len(mids), mids, err)
	t.Log(me.DeleteOtherFromChat("c6fac0ad4a214beaf6da808d6ba7e3025", ""))
}

func TestNewDailer(t *testing.T) {
	me := NewLogin()
	me.setDefaultHttpClient()
	t.Log(me.fast_connection.Get([]byte{}, "https://google.com"))
}

func TestCreateTempEmail(t *testing.T) {
	// t.Log(createTempEmail())
	// t.Log(getTempmailMessage("eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzUxMiJ9.eyJpYXQiOjE2MjMwNDk4NzAsInJvbGVzIjpbIlJPTEVfVVNFUiJdLCJ1c2VybmFtZSI6Ijg4MDhiNmQ1NjkwODFiZWQxMzhmOTlhMGIzMzg5MDRlQGxvZ2ljc3RyZWFrLmNvbSIsImlkIjoiNjBiZGM2OGQ0MTkyN2Y1YmI0MjBkZDkzIiwibWVyY3VyZSI6eyJzdWJzY3JpYmUiOlsiL2FjY291bnRzLzYwYmRjNjhkNDE5MjdmNWJiNDIwZGQ5MyJdfX0.kXHKiFxlV1QjpgZ9w8cszA7Xw4l2JSzl7Yte0ss6GK0rE4fmhg4iwqPveCrQCWlAXR2qxccV3MSGen-UqHTmhA"))
}
