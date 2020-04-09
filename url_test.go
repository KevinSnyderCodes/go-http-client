package http

import (
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

func TestRequest_FromURL(t *testing.T) {
	type args struct {
		u *url.URL
	}
	tests := []struct {
		name        string
		args        args
		wantRequest *Request
	}{
		{
			name: "success",
			args: args{
				u: &url.URL{
					Scheme:   "http",
					Host:     "www.test.com",
					Path:     "/api/v1/path",
					RawQuery: "foo=bar",
				},
			},
			wantRequest: &Request{
				Scheme: "http",
				Host:   "www.test.com",
				Path:   "/api/v1/path",
				Query: url.Values{
					"foo": []string{"bar"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Request{}
			o.FromURL(tt.args.u)
			if !reflect.DeepEqual(o, tt.wantRequest) {
				t.Errorf("Request = %v, want %v", o, tt.wantRequest)
			}
		})
	}
}

func TestRequest_FromURLString(t *testing.T) {
	type args struct {
		ref string
	}
	tests := []struct {
		name        string
		args        args
		wantErr     bool
		wantRequest *Request
	}{
		{
			name: "success",
			args: args{
				ref: "http://www.test.com/api/v1/path?foo=bar",
			},
			wantRequest: &Request{
				Scheme: "http",
				Host:   "www.test.com",
				Path:   "/api/v1/path",
				Query: url.Values{
					"foo": []string{"bar"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Request{}
			if err := o.FromURLString(tt.args.ref); (err != nil) != tt.wantErr {
				t.Errorf("Request.FromURLString() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(o, tt.wantRequest) {
				t.Errorf("Request = %v, want %v", o, tt.wantRequest)
			}
		})
	}
}

func TestRequest_URL(t *testing.T) {
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
		name    string
		fields  fields
		want    *url.URL
		wantErr bool
	}{
		{
			name: "success without query",
			fields: fields{
				Scheme: "http",
				Host:   "www.host.com",
				Path:   "/api/v1/path",
			},
			want: &url.URL{
				Scheme: "http",
				Host:   "www.host.com",
				Path:   "/api/v1/path",
			},
		},
		{
			name: "success with query",
			fields: fields{
				Scheme: "http",
				Host:   "www.host.com",
				Path:   "/api/v1/path",
				Query: url.Values{
					"foo": []string{"bar"},
				},
			},
			want: &url.URL{
				Scheme: "http",
				Host:   "www.host.com",
				Path:   "/api/v1/path",
				RawQuery: url.Values{
					"foo": []string{"bar"},
				}.Encode(),
			},
		},
		{
			name: "error no scheme",
			fields: fields{
				Host: "www.host.com",
				Path: "/api/v1/path",
			},
			wantErr: true,
		},
		{
			name: "error no host",
			fields: fields{
				Scheme: "http",
				Path:   "/api/v1/path",
			},
			wantErr: true,
		},
		{
			name: "error no path",
			fields: fields{
				Scheme: "http",
				Host:   "www.host.com",
			},
			wantErr: true,
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
			got, err := o.URL()
			if (err != nil) != tt.wantErr {
				t.Errorf("Request.URL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Request.URL() = %v, want %v", got, tt.want)
			}
		})
	}
}
