package endpoint

import (
	"io/github/pabloubal/ratelimiter/dto"
	"io/github/pabloubal/ratelimiter/internal/domain"
	"io/github/pabloubal/ratelimiter/internal/mapper"
)

type AddEndpointService interface {
	AddEndpoint(c dto.CreateEndpointDto) domain.RestApiError
}

type addEndpointService struct {
	limiter domain.Limiter
}

func NewAddEndpointUseCase(l domain.Limiter) *addEndpointService {
	return &addEndpointService{limiter: l}
}

func (this *addEndpointService) AddEndpoint(c dto.CreateEndpointDto) domain.RestApiError {
	endpoint, mapErr := mapper.CreateToEntity(c)

	if mapErr != nil {
		return mapErr
	}

	err := this.limiter.AddEndpoint(endpoint)

	return err
}
