package publicapi

import (
	"io/github/pabloubal/ratelimiter/infra/webapi"
	"io/github/pabloubal/ratelimiter/internal/mapper"
	"io/github/pabloubal/ratelimiter/internal/usecases/endpoint"
	"net/http"
)

type PublicEndpointsController interface {
	Request()
}

type publicEndpointsController struct {
	svcRQ endpoint.RequestEndpointService
}

func NewPublicEndpointsController(svcRQ endpoint.RequestEndpointService) *publicEndpointsController {
	return &publicEndpointsController{svcRQ: svcRQ}
}

func (e *publicEndpointsController) Request(w http.ResponseWriter, r *http.Request) {
	entity := mapper.RQToEntity(r)
	resp, err := e.svcRQ.RequestEndpoint(entity)
	if err != nil {
		webapi.SendError(err, w)
		return
	}

	webapi.SendResponse(resp, w)
}
