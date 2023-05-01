package domain

import (
	"fmt"
	"net/http"
	"strings"
)

type Request interface {
	Method() string
	Request() http.Request
	String() string
}

type request struct {
	rq *http.Request
}

func NewRequest(rq *http.Request) Request {
	return &request{rq: rq}
}

func (r *request) Method() string {
	return r.rq.Method
}

func (r *request) Request() http.Request {
	return *r.rq
}

func (r *request) String() string {
	path := strings.Split(r.rq.URL.Path, "/")[1]
	return fmt.Sprintf("%s %s", r.Method(), "/"+path)
}
