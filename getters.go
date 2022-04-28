package requests

import (
	"gopkg.in/headzoo/surf.v1"
	"io"
	"net/http"
)

func (r *Request) GetResponse() *http.Response {
	return r.response
}

func (r *Request) GetStatusCode() interface{} {
	return r.statusCode
}

func (r *Request) GetResponseBody() io.ReadCloser {
	return r.responseBody
}

func (r *Request) GetBody() []byte {
	return r.body
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
