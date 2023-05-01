package domain

import (
	"net/url"
	"time"
)

type RemoteEndpoint interface {
	IsLimitExceeded() bool
	IsTTLExceeded() bool
	Restart()
	Inc()
	Url() url.URL
}

type remoteEndpoint struct {
	path    string
	url     url.URL
	rqCount int
	max     int
	ttl     int64
}

func NewRemoteEndpoint(max int, urlObj url.URL, path string) (*remoteEndpoint, RestApiError) {
	if len(urlObj.Scheme) == 0 || len(urlObj.Host) == 0 || len(path) == 0 {
		return nil, ErrorBadURLFormat
	}
	return &remoteEndpoint{rqCount: 0, max: max, ttl: nextTTL(), url: urlObj, path: path}, nil
}

func nextTTL() int64 {
	return time.Now().Truncate(time.Minute).Add(time.Minute).UnixMilli()
}

func (l *remoteEndpoint) Restart() {
	l.ttl = nextTTL()
	l.rqCount = 0
}

func (l *remoteEndpoint) Inc() {
	l.rqCount++
}

func (l *remoteEndpoint) IsLimitExceeded() bool {
	return l.rqCount >= l.max
}

func (l *remoteEndpoint) IsTTLExceeded() bool {
	return l.ttl < time.Now().UnixMilli()
}

func (l *remoteEndpoint) Url() url.URL {
	return l.url
}
