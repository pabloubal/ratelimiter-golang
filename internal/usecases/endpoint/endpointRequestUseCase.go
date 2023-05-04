package endpoint

import (
	"io/github/pabloubal/ratelimiter/internal/domain"
	"net/http"
)

type RequestEndpointService interface {
	RequestEndpoint(entity domain.Request) (*http.Response, domain.RestApiError)
}

type requestEndpointService struct {
	limiter domain.Limiter
}

func NewRequestEndpointUseCase(l domain.Limiter) *requestEndpointService {
	return &requestEndpointService{limiter: l}
}

func (r *requestEndpointService) RequestEndpoint(entity domain.Request) (*http.Response, domain.RestApiError) {
	return r.limiter.RequestEndpoint(entity)
}
