package requests

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

type Request struct {
	BaseUrl      string
	EndPoint     string
	Headers      interface{}
	Parameters   map[interface{}]interface{}
	responseBody io.ReadCloser
	body         []byte
}

func (r *Request) Get() error {
	resp, err := http.Get(r.BaseUrl + r.EndPoint)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	r.responseBody = resp.Body

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	r.body = body
	return nil
}

func (r *Request) Post() error {
	postBody, _ := json.Marshal(r.Parameters)
	responseBody := bytes.NewBuffer(postBody)

	resp, err := http.Post(r.BaseUrl+r.EndPoint, "application/json", responseBody)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	r.responseBody = resp.Body

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	r.body = body

	return nil
}

func (r *Request) getResponseBody() io.ReadCloser {
	return r.responseBody
}

func (r *Request) getBody() []byte {
	return r.body
}
