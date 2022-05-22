package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type Request struct {
	BaseUrl      string
	EndPoint     string
	Headers      map[string]string
	Parameters   map[string]interface{}
	responseBody io.ReadCloser
	body         []byte
	requestBody  *[]byte
	response     *http.Response
	statusCode   interface{}
	Links        []string
}

func (r *Request) CloseResponseBody() {
	r.responseBody.Close()
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

func (r *Request) Delete() error {
	req, err := http.NewRequest("DELETE", r.BaseUrl+r.EndPoint, nil)
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
	var responseBody *bytes.Buffer
	if r.requestBody != nil {
		responseBody = bytes.NewBuffer(*r.requestBody)
	} else {
		postBody, err := json.Marshal(r.Parameters)
		if err != nil {
			fmt.Println(err)
		}
		responseBody = bytes.NewBuffer(postBody)
	}

	req, err := http.NewRequest("POST", r.BaseUrl+r.EndPoint, responseBody)
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
