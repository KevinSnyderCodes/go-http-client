package http

import (
	"net/http"
)

type Encoding string

const (
	EncodingUNKNOWN = ""
	EncodingJSON    = "JSON"
)

func inferEncoding(contentType string) Encoding {
	switch contentType {
	case "application/json":
		return EncodingJSON
	default:
		return EncodingUNKNOWN
	}
}

func (o *Request) inferRequestEncoding() Encoding {
	contentType := o.Header.Get("Content-Type")

	return inferEncoding(contentType)
}

func (o *Request) inferResponseEncoding(r *http.Response) Encoding {
	contentType := o.Header.Get("Accept")

	return inferEncoding(contentType)
}
