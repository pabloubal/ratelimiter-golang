package domain

import (
	"net/http"
	"strconv"
)

var (
	ErrorGenericRestAPI        = &genericRestAPIError{NewRestApiError("Generic error", http.StatusInternalServerError, nil)}
	ErrorEndpointAlreadyExists = &endpointAlreadyExistsError{NewRestApiError("Endpoint already registered", http.StatusBadRequest, nil)}
	ErrorEndpointDoesNotExist  = &endpointDoesNotExistError{NewRestApiError("Requested endpoint does not exist", http.StatusNotFound, nil)}
	ErrorBadURLFormat          = &badURLError{NewRestApiError("Bad URL Format", http.StatusBadRequest, nil)}
)

func ErrorLimitReached(limit int64) *endpointLimitReachedError {
	return &endpointLimitReachedError{
		NewRestApiError("Limit reached for requested resource",
			http.StatusTooManyRequests,
			http.Header{"Retry-After": {strconv.FormatInt(limit, 10)}},
		),
	}
}

func ErrorRemoteCall(msg string) *remoteCallError {
	return &remoteCallError{NewRestApiError(msg, http.StatusTooManyRequests, nil)}
}

/***********/

type RestApiError interface {
	Error() string
	StatusCode() int
	Header() http.Header
}

type restApiError struct {
	message    string
	statusCode int
	header     http.Header
}

func NewRestApiError(msg string, code int, header http.Header) *restApiError {
	return &restApiError{message: msg, statusCode: code, header: header}
}

func (r *restApiError) Error() string {
	return r.message
}

func (r *restApiError) StatusCode() int {
	return r.statusCode
}

func (r *restApiError) Header() http.Header {
	return r.header
}

type genericRestAPIError struct {
	*restApiError
}

type endpointAlreadyExistsError struct {
	*restApiError
}

type endpointDoesNotExistError struct {
	*restApiError
}

type endpointLimitReachedError struct {
	*restApiError
}

type remoteCallError struct {
	*restApiError
}

type badURLError struct {
	*restApiError
}
