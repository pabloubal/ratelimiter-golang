package endpoint

import (
	"io/github/pabloubal/ratelimiter/internal/domain"
	"io/github/pabloubal/ratelimiter/internal/mapper"
	"net/http"
)

type RequestEndpointService interface {
	RequestEndpoint(rq *http.Request) (*http.Response, domain.RestApiError)
}

type requestEndpointService struct {
	limiter domain.Limiter
}

func NewRequestEndpointUseCase(l domain.Limiter) *requestEndpointService {
	return &requestEndpointService{limiter: l}
}

func (r *requestEndpointService) RequestEndpoint(rq *http.Request) (*http.Response, domain.RestApiError) {
	entity := mapper.RQToEntity(rq)
	return r.limiter.RequestEndpoint(entity)
}
