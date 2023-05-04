package domain

import (
	"fmt"
	"io"
	"net/url"
	"strings"
)

type Request interface {
	Method() string
	String() string
	Body() io.ReadCloser
	Path() string
	RawQuery() string
	Header() map[string][]string
}

type request struct {
	method string
	urlObj url.URL
	body   io.ReadCloser
	header map[string][]string
}

func NewRequest(method string, urlObj url.URL, body io.ReadCloser, header map[string][]string) Request {
	return &request{method: method, urlObj: urlObj, body: body, header: header}
}

func (r *request) Method() string {
	return r.method
}

func (r *request) Body() io.ReadCloser {
	return r.body
}

func (r *request) Path() string {
	return r.urlObj.Path
}

func (r *request) RawQuery() string {
	return r.urlObj.RawQuery
}

func (r *request) Header() map[string][]string {
	return r.header
}

func (r *request) String() string {
	path := strings.Split(r.Path(), "/")[1]
	return fmt.Sprintf("%s %s", r.Method(), "/"+path)
}
