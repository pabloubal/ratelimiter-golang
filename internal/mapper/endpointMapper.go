package mapper

import (
	"io/github/pabloubal/ratelimiter/dto"
	"io/github/pabloubal/ratelimiter/internal/domain"
	"net/http"
)

func RQToEntity(rq *http.Request) domain.Request {
	return domain.NewRequest(rq)
}

func CreateToEntity(c dto.CreateEndpointDto) (domain.Endpoint, domain.RestApiError) {
	return domain.NewEndpoint(c.Path, c.Url, c.Limit, c.Method)
}

func CreateToDto(domain.Endpoint) dto.CreateEndpointDto {
	return dto.CreateEndpointDto{}
}
