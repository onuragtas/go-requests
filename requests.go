package go_requests

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Request struct {
	BaseUrl    string
	EndPoint   string
	Headers    interface{}
	Parameters map[interface{}]interface{}
}

func (r *Request) Get() ([]byte, error) {
	resp, err := http.Get(r.BaseUrl + r.EndPoint)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (r *Request) Post() ([]byte, error) {
	postBody, _ := json.Marshal(r.Parameters)
	responseBody := bytes.NewBuffer(postBody)

	resp, err := http.Post(r.BaseUrl+r.EndPoint, "application/json", responseBody)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
