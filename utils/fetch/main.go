package fetch

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Fetch struct {
	err        error
	ctype      string
	Body       []byte
	StringBody string
	header     http.Header
}

func Get(url string) *Fetch {
	fob := &Fetch{}
	resp, err := http.Get(url)
	if err != nil {
		fob.err = err
	}
	fob.header = resp.Header
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fob.err = err
	}
	fob.Body = body
	fob.StringBody = string(body)
	return fob
}

func (f *Fetch) Json(v interface{}) error {
	return json.Unmarshal(f.Body, &v)
}

func (f *Fetch) RawBody() []byte {
	return f.Body
}

func (f *Fetch) Header(v interface{}) http.Header {
	return f.header
}

func (f *Fetch) ContentType(ctype string) *Fetch {
	f.ctype = ctype
	return f
}
