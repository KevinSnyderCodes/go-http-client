package http

import (
	"net/http"
	"net/url"
	"reflect"
	"testing"
	"time"
)

type mockBody struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestRequest_Clear(t *testing.T) {
	type fields struct {
		Client       *http.Client
		Method       string
		Scheme       string
		Host         string
		Path         string
		Query        url.Values
		Header       http.Header
		RequestBody  interface{}
		ResponseBody interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   *Request
	}{
		{
			name: "success",
			fields: fields{
				Client: http.DefaultClient,
				Method: http.MethodGet,
				Scheme: "http",
				Host:   "www.example.com",
				Path:   "/api/v1/path",
				Query: url.Values{
					"foo": []string{"bar"},
				},
				Header: http.Header{
					"Foo": []string{"bar"},
				},
				RequestBody:  []byte("foo"),
				ResponseBody: &[]byte{},
			},
			want: &Request{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Request{
				Client:       tt.fields.Client,
				Method:       tt.fields.Method,
				Scheme:       tt.fields.Scheme,
				Host:         tt.fields.Host,
				Path:         tt.fields.Path,
				Query:        tt.fields.Query,
				Header:       tt.fields.Header,
				RequestBody:  tt.fields.RequestBody,
				ResponseBody: tt.fields.ResponseBody,
			}
			if got := o.Clear(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Request.Clear() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequest_WithClient(t *testing.T) {
	type fields struct {
		Client       *http.Client
		Method       string
		Scheme       string
		Host         string
		Path         string
		Query        url.Values
		Header       http.Header
		RequestBody  interface{}
		ResponseBody interface{}
	}
	type args struct {
		client *http.Client
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Request
	}{
		{
			name: "success empty client",
			args: args{
				client: &http.Client{},
			},
			want: &Request{
				Client: &http.Client{},
			},
		},
		{
			name: "success client with timeout",
			args: args{
				client: &http.Client{
					Timeout: 5 * time.Second,
				},
			},
			want: &Request{
				Client: &http.Client{
					Timeout: 5 * time.Second,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Request{
				Client:       tt.fields.Client,
				Method:       tt.fields.Method,
				Scheme:       tt.fields.Scheme,
				Host:         tt.fields.Host,
				Path:         tt.fields.Path,
				Query:        tt.fields.Query,
				Header:       tt.fields.Header,
				RequestBody:  tt.fields.RequestBody,
				ResponseBody: tt.fields.ResponseBody,
			}
			if got := o.WithClient(tt.args.client); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Request.WithClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequest_WithDefaultClient(t *testing.T) {
	type fields struct {
		Client       *http.Client
		Method       string
		Scheme       string
		Host         string
		Path         string
		Query        url.Values
		Header       http.Header
		RequestBody  interface{}
		ResponseBody interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   *Request
	}{
		{
			name: "success",
			want: &Request{
				Client: http.DefaultClient,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Request{
				Client:       tt.fields.Client,
				Method:       tt.fields.Method,
				Scheme:       tt.fields.Scheme,
				Host:         tt.fields.Host,
				Path:         tt.fields.Path,
				Query:        tt.fields.Query,
				Header:       tt.fields.Header,
				RequestBody:  tt.fields.RequestBody,
				ResponseBody: tt.fields.ResponseBody,
			}
			if got := o.WithDefaultClient(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Request.WithDefaultClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequest_WithTimeout(t *testing.T) {
	type fields struct {
		Client       *http.Client
		Method       string
		Scheme       string
		Host         string
		Path         string
		Query        url.Values
		Header       http.Header
		RequestBody  interface{}
		ResponseBody interface{}
	}
	type args struct {
		timeout time.Duration
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Request
	}{
		{
			name: "success without existing client",
			args: args{
				timeout: 5 * time.Second,
			},
			want: &Request{
				Client: &http.Client{
					Timeout: 5 * time.Second,
				},
			},
		},
		{
			name: "success with existing client",
			fields: fields{
				Client: &http.Client{
					Timeout: 1 * time.Second,
				},
			},
			args: args{
				timeout: 5 * time.Second,
			},
			want: &Request{
				Client: &http.Client{
					Timeout: 5 * time.Second,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Request{
				Client:       tt.fields.Client,
				Method:       tt.fields.Method,
				Scheme:       tt.fields.Scheme,
				Host:         tt.fields.Host,
				Path:         tt.fields.Path,
				Query:        tt.fields.Query,
				Header:       tt.fields.Header,
				RequestBody:  tt.fields.RequestBody,
				ResponseBody: tt.fields.ResponseBody,
			}
			if got := o.WithTimeout(tt.args.timeout); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Request.WithTimeout() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequest_WithMethod(t *testing.T) {
	type fields struct {
		Client       *http.Client
		Method       string
		Scheme       string
		Host         string
		Path         string
		Query        url.Values
		Header       http.Header
		RequestBody  interface{}
		ResponseBody interface{}
	}
	type args struct {
		method string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Request
	}{
		{
			name: "success raw string",
			args: args{
				method: "foo",
			},
			want: &Request{
				Method: "foo",
			},
		},
		{
			name: "success http method string",
			args: args{
				method: http.MethodGet,
			},
			want: &Request{
				Method: http.MethodGet,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Request{
				Client:       tt.fields.Client,
				Method:       tt.fields.Method,
				Scheme:       tt.fields.Scheme,
				Host:         tt.fields.Host,
				Path:         tt.fields.Path,
				Query:        tt.fields.Query,
				Header:       tt.fields.Header,
				RequestBody:  tt.fields.RequestBody,
				ResponseBody: tt.fields.ResponseBody,
			}
			if got := o.WithMethod(tt.args.method); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Request.WithMethod() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequest_WithScheme(t *testing.T) {
	type fields struct {
		Client       *http.Client
		Method       string
		Scheme       string
		Host         string
		Path         string
		Query        url.Values
		Header       http.Header
		RequestBody  interface{}
		ResponseBody interface{}
	}
	type args struct {
		scheme string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Request
	}{
		{
			name: "success",
			args: args{
				scheme: "http",
			},
			want: &Request{
				Scheme: "http",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Request{
				Client:       tt.fields.Client,
				Method:       tt.fields.Method,
				Scheme:       tt.fields.Scheme,
				Host:         tt.fields.Host,
				Path:         tt.fields.Path,
				Query:        tt.fields.Query,
				Header:       tt.fields.Header,
				RequestBody:  tt.fields.RequestBody,
				ResponseBody: tt.fields.ResponseBody,
			}
			if got := o.WithScheme(tt.args.scheme); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Request.WithScheme() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequest_WithHost(t *testing.T) {
	type fields struct {
		Client       *http.Client
		Method       string
		Scheme       string
		Host         string
		Path         string
		Query        url.Values
		Header       http.Header
		RequestBody  interface{}
		ResponseBody interface{}
	}
	type args struct {
		host string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Request
	}{
		{
			name: "success",
			args: args{
				host: "www.foo.com",
			},
			want: &Request{
				Host: "www.foo.com",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Request{
				Client:       tt.fields.Client,
				Method:       tt.fields.Method,
				Scheme:       tt.fields.Scheme,
				Host:         tt.fields.Host,
				Path:         tt.fields.Path,
				Query:        tt.fields.Query,
				Header:       tt.fields.Header,
				RequestBody:  tt.fields.RequestBody,
				ResponseBody: tt.fields.ResponseBody,
			}
			if got := o.WithHost(tt.args.host); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Request.WithHost() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequest_WithPath(t *testing.T) {
	type fields struct {
		Client       *http.Client
		Method       string
		Scheme       string
		Host         string
		Path         string
		Query        url.Values
		Header       http.Header
		RequestBody  interface{}
		ResponseBody interface{}
	}
	type args struct {
		path string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Request
	}{
		{
			name: "success",
			args: args{
				path: "/api/v1/foo",
			},
			want: &Request{
				Path: "/api/v1/foo",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Request{
				Client:       tt.fields.Client,
				Method:       tt.fields.Method,
				Scheme:       tt.fields.Scheme,
				Host:         tt.fields.Host,
				Path:         tt.fields.Path,
				Query:        tt.fields.Query,
				Header:       tt.fields.Header,
				RequestBody:  tt.fields.RequestBody,
				ResponseBody: tt.fields.ResponseBody,
			}
			if got := o.WithPath(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Request.WithPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequest_WithQuery(t *testing.T) {
	type fields struct {
		Client       *http.Client
		Method       string
		Scheme       string
		Host         string
		Path         string
		Query        url.Values
		Header       http.Header
		RequestBody  interface{}
		ResponseBody interface{}
	}
	type args struct {
		query url.Values
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Request
	}{
		{
			name: "success without existing query",
			args: args{
				query: url.Values{
					"foo": []string{"bar"},
				},
			},
			want: &Request{
				Query: url.Values{
					"foo": []string{"bar"},
				},
			},
		},
		{
			name: "success with existing query",
			fields: fields{
				Query: url.Values{
					"bar": []string{"baz"},
				},
			},
			args: args{
				query: url.Values{
					"foo": []string{"bar"},
				},
			},
			want: &Request{
				Query: url.Values{
					"foo": []string{"bar"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Request{
				Client:       tt.fields.Client,
				Method:       tt.fields.Method,
				Scheme:       tt.fields.Scheme,
				Host:         tt.fields.Host,
				Path:         tt.fields.Path,
				Query:        tt.fields.Query,
				Header:       tt.fields.Header,
				RequestBody:  tt.fields.RequestBody,
				ResponseBody: tt.fields.ResponseBody,
			}
			if got := o.WithQuery(tt.args.query); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Request.WithQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequest_WithDefaultQuery(t *testing.T) {
	type fields struct {
		Client       *http.Client
		Method       string
		Scheme       string
		Host         string
		Path         string
		Query        url.Values
		Header       http.Header
		RequestBody  interface{}
		ResponseBody interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   *Request
	}{
		{
			name: "success",
			want: &Request{
				Query: url.Values{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Request{
				Client:       tt.fields.Client,
				Method:       tt.fields.Method,
				Scheme:       tt.fields.Scheme,
				Host:         tt.fields.Host,
				Path:         tt.fields.Path,
				Query:        tt.fields.Query,
				Header:       tt.fields.Header,
				RequestBody:  tt.fields.RequestBody,
				ResponseBody: tt.fields.ResponseBody,
			}
			if got := o.WithDefaultQuery(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Request.WithDefaultQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequest_AddQuery(t *testing.T) {
	type fields struct {
		Client       *http.Client
		Method       string
		Scheme       string
		Host         string
		Path         string
		Query        url.Values
		Header       http.Header
		RequestBody  interface{}
		ResponseBody interface{}
	}
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Request
	}{
		{
			name: "success without existing query",
			args: args{
				key:   "foo",
				value: "bar",
			},
			want: &Request{
				Query: url.Values{
					"foo": []string{"bar"},
				},
			},
		},
		{
			name: "success with existing query without conflict",
			fields: fields{
				Query: url.Values{
					"bar": []string{"baz"},
				},
			},
			args: args{
				key:   "foo",
				value: "bar",
			},
			want: &Request{
				Query: url.Values{
					"foo": []string{"bar"},
					"bar": []string{"baz"},
				},
			},
		},
		{
			name: "success with existing query with conflict different value",
			fields: fields{
				Query: url.Values{
					"foo": []string{"bar"},
				},
			},
			args: args{
				key:   "foo",
				value: "baz",
			},
			want: &Request{
				Query: url.Values{
					"foo": []string{"bar", "baz"},
				},
			},
		},
		{
			name: "success with existing query with conflict same value",
			fields: fields{
				Query: url.Values{
					"foo": []string{"bar"},
				},
			},
			args: args{
				key:   "foo",
				value: "bar",
			},
			want: &Request{
				Query: url.Values{
					"foo": []string{"bar", "bar"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Request{
				Client:       tt.fields.Client,
				Method:       tt.fields.Method,
				Scheme:       tt.fields.Scheme,
				Host:         tt.fields.Host,
				Path:         tt.fields.Path,
				Query:        tt.fields.Query,
				Header:       tt.fields.Header,
				RequestBody:  tt.fields.RequestBody,
				ResponseBody: tt.fields.ResponseBody,
			}
			if got := o.AddQuery(tt.args.key, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Request.AddQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequest_WithHeader(t *testing.T) {
	type fields struct {
		Client       *http.Client
		Method       string
		Scheme       string
		Host         string
		Path         string
		Query        url.Values
		Header       http.Header
		RequestBody  interface{}
		ResponseBody interface{}
	}
	type args struct {
		header http.Header
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Request
	}{
		{
			name: "success without existing header",
			args: args{
				header: http.Header{
					"foo": []string{"bar"},
				},
			},
			want: &Request{
				Header: http.Header{
					"foo": []string{"bar"},
				},
			},
		},
		{
			name: "success with existing header",
			fields: fields{
				Header: http.Header{
					"bar": []string{"baz"},
				},
			},
			args: args{
				header: http.Header{
					"foo": []string{"bar"},
				},
			},
			want: &Request{
				Header: http.Header{
					"foo": []string{"bar"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Request{
				Client:       tt.fields.Client,
				Method:       tt.fields.Method,
				Scheme:       tt.fields.Scheme,
				Host:         tt.fields.Host,
				Path:         tt.fields.Path,
				Query:        tt.fields.Query,
				Header:       tt.fields.Header,
				RequestBody:  tt.fields.RequestBody,
				ResponseBody: tt.fields.ResponseBody,
			}
			if got := o.WithHeader(tt.args.header); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Request.WithHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequest_WithDefaultHeader(t *testing.T) {
	type fields struct {
		Client       *http.Client
		Method       string
		Scheme       string
		Host         string
		Path         string
		Query        url.Values
		Header       http.Header
		RequestBody  interface{}
		ResponseBody interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   *Request
	}{
		{
			name: "success",
			want: &Request{
				Header: http.Header{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Request{
				Client:       tt.fields.Client,
				Method:       tt.fields.Method,
				Scheme:       tt.fields.Scheme,
				Host:         tt.fields.Host,
				Path:         tt.fields.Path,
				Query:        tt.fields.Query,
				Header:       tt.fields.Header,
				RequestBody:  tt.fields.RequestBody,
				ResponseBody: tt.fields.ResponseBody,
			}
			if got := o.WithDefaultHeader(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Request.WithDefaultHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequest_AddHeader(t *testing.T) {
	type fields struct {
		Client       *http.Client
		Method       string
		Scheme       string
		Host         string
		Path         string
		Query        url.Values
		Header       http.Header
		RequestBody  interface{}
		ResponseBody interface{}
	}
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Request
	}{
		{
			name: "success without existing header",
			args: args{
				key:   "foo",
				value: "bar",
			},
			want: &Request{
				Header: http.Header{
					"Foo": []string{"bar"},
				},
			},
		},
		{
			name: "success with existing header without conflict",
			fields: fields{
				Header: http.Header{
					"Bar": []string{"baz"},
				},
			},
			args: args{
				key:   "foo",
				value: "bar",
			},
			want: &Request{
				Header: http.Header{
					"Foo": []string{"bar"},
					"Bar": []string{"baz"},
				},
			},
		},
		{
			name: "success with existing header with conflict different value",
			fields: fields{
				Header: http.Header{
					"Foo": []string{"bar"},
				},
			},
			args: args{
				key:   "foo",
				value: "baz",
			},
			want: &Request{
				Header: http.Header{
					"Foo": []string{"bar", "baz"},
				},
			},
		},
		{
			name: "success with existing header with conflict same value",
			fields: fields{
				Header: http.Header{
					"Foo": []string{"bar"},
				},
			},
			args: args{
				key:   "foo",
				value: "bar",
			},
			want: &Request{
				Header: http.Header{
					"Foo": []string{"bar", "bar"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Request{
				Client:       tt.fields.Client,
				Method:       tt.fields.Method,
				Scheme:       tt.fields.Scheme,
				Host:         tt.fields.Host,
				Path:         tt.fields.Path,
				Query:        tt.fields.Query,
				Header:       tt.fields.Header,
				RequestBody:  tt.fields.RequestBody,
				ResponseBody: tt.fields.ResponseBody,
			}
			if got := o.AddHeader(tt.args.key, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Request.AddHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewRequest(t *testing.T) {
	tests := []struct {
		name string
		want *Request
	}{
		{
			name: "success",
			want: &Request{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRequest(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequest_WithRequestBody(t *testing.T) {
	type fields struct {
		Client       *http.Client
		Method       string
		Scheme       string
		Host         string
		Path         string
		Query        url.Values
		Header       http.Header
		RequestBody  interface{}
		ResponseBody interface{}
	}
	type args struct {
		requestBody interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Request
	}{
		{
			name: "success []byte",
			args: args{
				requestBody: []byte("hello"),
			},
			want: &Request{
				RequestBody: []byte("hello"),
			},
		},
		{
			name: "success mock body",
			args: args{
				requestBody: &mockBody{},
			},
			want: &Request{
				RequestBody: &mockBody{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Request{
				Client:       tt.fields.Client,
				Method:       tt.fields.Method,
				Scheme:       tt.fields.Scheme,
				Host:         tt.fields.Host,
				Path:         tt.fields.Path,
				Query:        tt.fields.Query,
				Header:       tt.fields.Header,
				RequestBody:  tt.fields.RequestBody,
				ResponseBody: tt.fields.ResponseBody,
			}
			if got := o.WithRequestBody(tt.args.requestBody); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Request.WithRequestBody() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequest_WithResponseBody(t *testing.T) {
	type fields struct {
		Client       *http.Client
		Method       string
		Scheme       string
		Host         string
		Path         string
		Query        url.Values
		Header       http.Header
		RequestBody  interface{}
		ResponseBody interface{}
	}
	type args struct {
		responseBody interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Request
	}{
		{
			name: "success mock body",
			args: args{
				responseBody: &mockBody{},
			},
			want: &Request{
				ResponseBody: &mockBody{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Request{
				Client:       tt.fields.Client,
				Method:       tt.fields.Method,
				Scheme:       tt.fields.Scheme,
				Host:         tt.fields.Host,
				Path:         tt.fields.Path,
				Query:        tt.fields.Query,
				Header:       tt.fields.Header,
				RequestBody:  tt.fields.RequestBody,
				ResponseBody: tt.fields.ResponseBody,
			}
			if got := o.WithResponseBody(tt.args.responseBody); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Request.WithResponseBody() = %v, want %v", got, tt.want)
			}
		})
	}
}
