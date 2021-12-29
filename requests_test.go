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

	request := Request{BaseUrl: "https://www.akakce.com/c/?z=124&v=11281&s=1&p=1094026950&c=15850&k=453&g=1593405445&f=%2Fr%2F%3Fpr%3D1094026950%26vd%3D11281%26pg%3D1593405445", EndPoint: "", Headers: headers}
	err := request.GetWithJS()
	if err != nil {
		t.Fatalf("Error %s", err.Error())
	}

	if request.response.StatusCode != 200 {
		t.Fatalf(`%q, %v, want match for %#q, nil`, request.response.StatusCode, err, 200)
	}
}
