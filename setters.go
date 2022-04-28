package requests

func (r *Request) SetRequestBody(body *[]byte) {
	r.requestBody = body
}
