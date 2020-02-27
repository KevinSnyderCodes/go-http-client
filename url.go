package http

import (
	"fmt"
	"net/url"
)

// URL builds a URL object from the Request.
func (o *Request) URL() (*url.URL, error) {
	if o.Scheme == "" {
		return nil, fmt.Errorf("must provide scheme")
	}
	if o.Host == "" {
		return nil, fmt.Errorf("must provide host")
	}
	if o.Path == "" {
		return nil, fmt.Errorf("must provide path")
	}

	return &url.URL{
		Scheme:   o.Scheme,
		Host:     o.Host,
		Path:     o.Path,
		RawQuery: o.Query.Encode(),
	}, nil
}
