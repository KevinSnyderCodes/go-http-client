package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// DoOptions are options to use when making a HTTP request.
type DoOptions struct {
	WithRequestEncoding  Encoding
	WithResponseEncoding Encoding
}

func joinOptions(options ...*DoOptions) *DoOptions {
	opts := DoOptions{}

	for _, o := range options {
		if opts.WithRequestEncoding == "" {
			opts.WithRequestEncoding = o.WithRequestEncoding
		}
		if opts.WithResponseEncoding == "" {
			opts.WithResponseEncoding = o.WithResponseEncoding
		}
	}

	return &opts
}

// Do makes the HTTP request.
func (o *Request) Do(options ...*DoOptions) (*http.Response, error) {
	opts := joinOptions(options...)

	o.ensure()

	if o.Method == "" {
		return nil, fmt.Errorf("must provide method")
	}

	u, err := o.URL()
	if err != nil {
		return nil, fmt.Errorf("error building URL: %w", err)
	}

	var reqBody []byte
	if o.RequestBody != nil {
		switch v := o.RequestBody.(type) {
		case []byte:
			reqBody = v
		default:
			encoding := opts.WithRequestEncoding
			if encoding == "" {
				encoding = o.inferRequestEncoding()
			}

			switch encoding {
			case EncodingJSON:
				var err error
				reqBody, err = json.Marshal(o.RequestBody)
				if err != nil {
					return nil, fmt.Errorf("error encoding request body: %w", err)
				}
			default:
				return nil, fmt.Errorf("unable to encode request body")
			}
		}
	}

	req, err := http.NewRequest(o.Method, u.String(), bytes.NewBuffer(reqBody))
	req.Header = o.Header
	if err != nil {
		return nil, fmt.Errorf("error creating http request: %w", err)
	}

	resp, err := o.Client.Do(req)
	if err != nil {
		return resp, fmt.Errorf("error making http request: %w", err)
	}
	if resp.StatusCode/100 > 2 {
		err = &StatusCodeError{StatusCode: resp.StatusCode}
	}

	if o.ResponseBody != nil {
		encoding := opts.WithResponseEncoding
		if encoding == "" {
			o.inferResponseEncoding(resp)
		}

		switch encoding {
		case EncodingJSON:
			if err := json.NewDecoder(resp.Body).Decode(o.ResponseBody); err != nil {
				return resp, fmt.Errorf("error decoding response body: %w", err)
			}
		default:
			return resp, fmt.Errorf("unable to decode request body")
		}
	}

	return resp, err
}
