package domain

import (
	"log"
	"net/http"
	"sync"
)

type Limiter interface {
	AddEndpoint(e Endpoint) RestApiError
	RequestEndpoint(r Request) (*http.Response, RestApiError)
}

type limiter struct {
	limMtx          sync.Mutex
	remoteEndpoints map[string]*remoteEndpoint
}

func NewLimiter() *limiter {
	var l limiter
	l.remoteEndpoints = make(map[string]*remoteEndpoint)
	return &l
}

func (l *limiter) AddEndpoint(e Endpoint) RestApiError {
	_, exists := l.remoteEndpoints[e.String()]

	if exists {
		return ErrorEndpointAlreadyExists
	}

	l.limMtx.Lock()
	defer l.limMtx.Unlock()
	re, err := NewRemoteEndpoint(e.GetLimit(), e.GetUrl(), e.GetPath())
	if err != nil {
		return err
	}

	l.remoteEndpoints[e.String()] = re

	return nil
}

func (l *limiter) RequestEndpoint(rq Request) (*http.Response, RestApiError) {
	remote, exists := l.remoteEndpoints[rq.String()]

	if !exists {
		return nil, ErrorEndpointDoesNotExist
	}

	if !l.isBelowLimit(remote) {
		return nil, ErrorLimitReached(remote.ttl)
	}

	response, err := l.callEndpoint(remote, rq)

	return response, err
}

func (l *limiter) isBelowLimit(rm RemoteEndpoint) bool {
	l.limMtx.Lock()
	defer l.limMtx.Unlock()

	if rm.IsLimitExceeded() {
		if rm.IsTTLExceeded() {
			rm.Restart()
		} else {
			return false
		}
	}

	rm.Inc()
	return true
}

func (l *limiter) callEndpoint(rm RemoteEndpoint, rq Request) (*http.Response, RestApiError) {
	client := &http.Client{}

	url := GetUrl(rm.Url(), rq)

	request, err := http.NewRequest(rq.Method(), url.String(), rq.Body())
	defer rq.Body().Close()
	if err != nil {
		log.Println(err)
		return nil, ErrorRemoteCall(err.Error())
	}
	request.Header = rq.Header()

	resp, err := client.Do(request)
	if err != nil {
		log.Println(err)
		return nil, ErrorRemoteCall(err.Error())
	}

	return resp, nil
}
