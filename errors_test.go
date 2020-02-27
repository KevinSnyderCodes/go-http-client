package http

import (
	"net/http"
	"testing"
)

func TestStatusCodeError_Error(t *testing.T) {
	type fields struct {
		StatusCode int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "success raw int",
			fields: fields{
				StatusCode: 123,
			},
			want: "received status code 123",
		},
		{
			name: "success http status code int",
			fields: fields{
				StatusCode: http.StatusNotFound,
			},
			want: "received status code 404",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &StatusCodeError{
				StatusCode: tt.fields.StatusCode,
			}
			if got := o.Error(); got != tt.want {
				t.Errorf("StatusCodeError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}
