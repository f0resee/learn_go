package builder

type Request struct {
	Method string
	Header map[string]string
	Proto  string
}

func NewRequestBuilder() *RequestBuilder {
	return &RequestBuilder{
		request: &Request{},
	}
}

type RequestBuilder struct {
	request *Request
}

func (r *RequestBuilder) Method(method string) *RequestBuilder {
	r.request.Method = method
	return r
}

func (r *RequestBuilder) Header(headers map[string]string) *RequestBuilder {
	r.request.Header = headers
	return r
}

func (r *RequestBuilder) Proto(proto string) *RequestBuilder {
	r.request.Proto = proto
	return r
}

func (r *RequestBuilder) Build() *Request {
	return r.request
}
