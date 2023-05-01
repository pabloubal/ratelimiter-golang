package webapi

import (
	"encoding/json"
	"io"
	"io/github/pabloubal/ratelimiter/internal/domain"
	"net/http"
)

func SendResponse(r *http.Response, w http.ResponseWriter) {
	defer r.Body.Close()
	copyHeaders(w.Header(), r.Header)
	w.WriteHeader(r.StatusCode)
	io.Copy(w, r.Body)
}

func SendError(err domain.RestApiError, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	copyHeaders(w.Header(), err.Header())
	w.WriteHeader(err.StatusCode())
	return json.NewEncoder(w).Encode(err.Error())
}

func copyHeaders(dest, src http.Header) {
	if src == nil {
		return
	}

	for k, vs := range src {
		for _, v := range vs {
			dest.Add(k, v)
		}
	}
}
