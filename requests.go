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
	Headers      map[string]string
	Parameters   map[interface{}]interface{}
	responseBody io.ReadCloser
	body         []byte
	response     *http.Response
}

func (r *Request) Get() error {
	req, err := http.NewRequest("GET", r.BaseUrl+r.EndPoint, nil)
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}
	for key, value := range r.Headers {
		req.Header.Set(key, value)
	}

	resp, err := http.DefaultClient.Do(req)
	r.response = resp

	if err != nil {
		return err
	}

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
	r.response = resp

	if err != nil {
		return err
	}

	r.responseBody = resp.Body

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	r.body = body

	return nil
}

func (r *Request) GetResponseBody() io.ReadCloser {
	return r.responseBody
}

func (r *Request) GetBody() []byte {
	return r.body
}

func (r *Request) CloseResponseBody() {
	r.responseBody.Close()
}
