package requests

import (
	"bytes"
	"encoding/json"
	"gopkg.in/headzoo/surf.v1"
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
	statusCode   interface{}
	Links        []string
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

func (r *Request) GetWithJS() error {
	var links []string
	bow := surf.NewBrowser()
	err := bow.Open(r.BaseUrl + r.EndPoint)
	if err != nil {
		panic(err)
	}

	if bow.StatusCode() == 403 {
		bow.Reload()
	}

	for _, link := range bow.Links() {
		links = append(links, link.Url().String())
	}

	r.Links = links

	r.statusCode = bow.StatusCode()

	r.body = []byte(bow.Body())

	return nil

}
