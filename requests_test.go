package requests

import (
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestGet(t *testing.T) {
	headers := make(map[string]string)

	headers["Content-Type"] = "application/json"
	headers["User-Agent"] = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.103 Safari/537.36"

	request := Request{BaseUrl: "http://google.com/", EndPoint: "", Headers: headers}
	err := request.Get()
	if err != nil {
		t.Fatalf("Error %s", err.Error())
	}

	if request.response.StatusCode != 200 {
		t.Fatalf(`%q, %v, want match for %#q, nil`, request.response.StatusCode, err, 200)
	}
}
