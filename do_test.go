package http

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
)

func TestRequest_Do(t *testing.T) {
	type fields struct {
		Client *http.Client
		Method string
		// Scheme       string
		// Host         string
		Path         string
		Query        url.Values
		Header       http.Header
		RequestBody  interface{}
		ResponseBody interface{}
	}
	type args struct {
		options []*DoOptions
	}
	type server struct {
		wantBody []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		server  server
		want    *http.Response
		wantErr bool
	}{
		{
			name: "success get",
			fields: fields{
				Method: http.MethodGet,
				Path:   "/api/v1/path",
			},
		},
		{
			name: "success get with header",
			fields: fields{
				Method: http.MethodGet,
				Path:   "/api/v1/path",
				Header: http.Header{
					"Foo": []string{"bar"},
				},
			},
		},
		{
			name: "success get with query",
			fields: fields{
				Method: http.MethodGet,
				Path:   "/api/v1/path",
				Query: url.Values{
					"foo": []string{"bar"},
				},
			},
		},
		{
			name: "success post",
			fields: fields{
				Method: http.MethodPost,
				Path:   "/api/v1/path",
			},
		},
		{
			name: "success post with body",
			fields: fields{
				Method:      http.MethodPost,
				Path:        "/api/v1/path",
				RequestBody: []byte("foo"),
			},
			server: server{
				wantBody: []byte("foo"),
			},
		},
		{
			name: "error no method",
			fields: fields{
				Path: "/api/v1/path",
			},
			wantErr: true,
		},
		{
			name: "error no path",
			fields: fields{
				Method: http.MethodGet,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		// Ensure that wanted values have defaults
		if tt.server.wantBody == nil {
			tt.server.wantBody = []byte{}
		}

		// Create test server
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if !reflect.DeepEqual(r.Method, tt.fields.Method) {
					t.Errorf("http.Request.Method = %v, want %v", r.Method, tt.fields.Method)
				}
				for key := range tt.fields.Header {
					if !reflect.DeepEqual(r.Header.Get(key), tt.fields.Header.Get(key)) {
						t.Errorf("http.Request.Header.Get(\"%s\") = %v, want %v", key, r.Header.Get(key), tt.fields.Header.Get(key))
					}
				}
				if !reflect.DeepEqual(r.URL.Path, tt.fields.Path) {
					t.Errorf("http.Request.URL.Path = %v, want %v", r.URL.Path, tt.fields.Path)
				}
				for key := range tt.fields.Query {
					if !reflect.DeepEqual(r.URL.Query().Get(key), tt.fields.Query.Get(key)) {
						t.Errorf("http.Request.URL.Query().Get(\"%s\") = %v, want %v", key, r.URL.Query().Get(key), tt.fields.Query.Get(key))
					}
				}

				body, err := ioutil.ReadAll(r.Body)
				if err != nil {
					t.Fatal(err)
				}
				if !reflect.DeepEqual(body, tt.server.wantBody) {
					t.Errorf("ioutil.ReadAll(http.Request.Body) = %v, want %v", body, tt.server.wantBody)
				}
			}))

			// Parse URL from test server
			u, err := url.Parse(server.URL)
			if err != nil {
				t.Fatal(err)
			}

			// Build and make request
			o := &Request{
				Client:       tt.fields.Client,
				Method:       tt.fields.Method,
				Scheme:       u.Scheme,
				Host:         u.Host,
				Path:         tt.fields.Path,
				Query:        tt.fields.Query,
				Header:       tt.fields.Header,
				RequestBody:  tt.fields.RequestBody,
				ResponseBody: tt.fields.ResponseBody,
			}
			_, err = o.Do(tt.args.options...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Request.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("Request.Do() = %v, want %v", got, tt.want)
			// }

			server.Close()
		})
	}
}
