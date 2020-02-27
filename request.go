package http

import (
	"net/http"
	"net/url"
	"time"
)

// Request is a HTTP request.
type Request struct {
	Client *http.Client

	Method       string
	Scheme       string
	Host         string
	Path         string
	Query        url.Values
	Header       http.Header
	RequestBody  interface{}
	ResponseBody interface{}
}

// WithClient sets the HTTP client of the Request.
func (o *Request) WithClient(client *http.Client) *Request {
	o.Client = client
	return o
}

// WithDefaultClient sets the HTTP client of the Request to the default.
func (o *Request) WithDefaultClient() *Request {
	o.Client = http.DefaultClient
	return o
}

// WithTimeout sets the HTTP client timeout of the request.
func (o *Request) WithTimeout(timeout time.Duration) *Request {
	o.ensureClient()

	o.Client.Timeout = timeout
	return o
}

// WithMethod sets the method of the Request.
func (o *Request) WithMethod(method string) *Request {
	o.Method = method
	return o
}

// WithScheme sets the scheme of the Request.
func (o *Request) WithScheme(scheme string) *Request {
	o.Scheme = scheme
	return o
}

// WithHost sets the host of the Request.
func (o *Request) WithHost(host string) *Request {
	o.Host = host
	return o
}

// WithPath sets the path of the Request.
func (o *Request) WithPath(path string) *Request {
	o.Path = path
	return o
}

// WithQuery sets the query of the Request.
func (o *Request) WithQuery(query url.Values) *Request {
	o.Query = query
	return o
}

// WithDefaultQuery sets the query of the Request to the default.
func (o *Request) WithDefaultQuery() *Request {
	o.WithQuery(url.Values{})
	return o
}

// AddQuery adds a key-value pair to the query of the Request.
func (o *Request) AddQuery(key, value string) *Request {
	o.ensureQuery()

	o.Query.Add(key, value)
	return o
}

// WithHeader sets the header of the Request.
func (o *Request) WithHeader(header http.Header) *Request {
	o.Header = header
	return o
}

// WithDefaultHeader sets the header of the Request to the default.
func (o *Request) WithDefaultHeader() *Request {
	o.WithHeader(http.Header{})
	return o
}

// AddHeader adds a key-value pair to the header of the Request.
func (o *Request) AddHeader(key, value string) *Request {
	o.ensureHeader()

	o.Header.Add(key, value)
	return o
}

// WithRequestBody sets the request body of the Request.
//
// If the request body is a []byte then it will be used as the exact value of
// the request body. All other types will be encoded according to the
// "Content-Type" specified in the request header.
func (o *Request) WithRequestBody(requestBody interface{}) *Request {
	o.RequestBody = requestBody
	return o
}

// WithResponseBody sets the response body of the Request.
//
// The response body will be decoded according to the "Content-Type" specified
// in the response header, or the "Accept" specified in the request body if the
// former is not specified.
func (o *Request) WithResponseBody(responseBody interface{}) *Request {
	o.ResponseBody = responseBody
	return o
}

// NewRequest creates a new Request.
func NewRequest() *Request {
	return &Request{}
}
