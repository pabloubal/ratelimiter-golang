package domain

import (
	"net/url"
	"strings"
)

func GetUrl(urlObj url.URL, rq Request) url.URL {
	host := urlObj.Host
	scheme := urlObj.Scheme
	if len(scheme) == 0 {
		scheme = "http"
	}
	rqPath := rq.Request().URL.Path
	var path string

	if len(urlObj.Path) > 0 {
		path = strings.TrimSuffix(urlObj.Path+"/"+strings.Replace(rqPath, urlObj.Path, "", 1), "/")
	} else {
		path = rqPath
	}

	query := rq.Request().URL.RawQuery

	return url.URL{Scheme: scheme, Host: host, Path: path, RawQuery: query}
}
