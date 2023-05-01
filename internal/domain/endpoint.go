package domain

import (
	"fmt"
	"net/http"
	"net/url"
)

type Endpoint interface {
	GetUrl() url.URL
	GetMethod() string
	GetLimit() int
	String() string
	GetPath() string
}

type endpoint struct {
	Path    string
	Url     url.URL
	Limit   int
	Method  string
	Request *http.Request
}

func NewEndpoint(path string, urlStr string, limit int, method string) (Endpoint, RestApiError) {
	urlObj, err := url.Parse(urlStr)
	if err != nil {
		return nil, ErrorGenericRestAPI
	}
	return &endpoint{
		Path:   path,
		Url:    *urlObj,
		Limit:  limit,
		Method: method,
	}, nil
}

func (e *endpoint) String() string {
	return fmt.Sprintf("%s %s", e.Method, e.Path)
}

func (e *endpoint) GetUrl() url.URL {
	return e.Url
}

func (e *endpoint) GetMethod() string {
	return e.Method
}

func (e *endpoint) GetLimit() int {
	return e.Limit
}

func (e *endpoint) GetPath() string {
	return e.Path
}
