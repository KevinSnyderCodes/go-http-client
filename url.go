package http

import (
	"fmt"
	"net/url"
)

// FromURL sets the scheme, host, path, and query of the Request.
func (o *Request) FromURL(u *url.URL) *Request {
	if u.Scheme != "" {
		o.WithScheme(u.Scheme)
	}
	if u.Host != "" {
		o.WithHost(u.Host)
	}
	if u.Path != "" {
		o.WithPath(u.Path)
	}
	if len(u.Query()) > 0 {
		o.WithQuery(u.Query())
	}

	return o
}

// FromURLString sets the scheme, host, path, and query of the Request.
func (o *Request) FromURLString(ref string) (*Request, error) {
	u := &url.URL{}
	u, err := u.Parse(ref)
	if err != nil {
		return nil, fmt.Errorf("error parsing URL: %w", err)
	}

	return o.FromURL(u), nil
}

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
