package todos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
)

// FindURL generates an URL for the find operation
type FindURL struct {
}

// Build a url path and query string
func (o *FindURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/"

	result.Path = _path

	return &result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *FindURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *FindURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *FindURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on FindURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on FindURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *FindURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
