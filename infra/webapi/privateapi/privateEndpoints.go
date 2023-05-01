package privateapi

import (
	"encoding/json"
	"io/github/pabloubal/ratelimiter/dto"
	"io/github/pabloubal/ratelimiter/infra/webapi"
	"io/github/pabloubal/ratelimiter/internal/usecases/endpoint"
	"net/http"
)

type PrivateEndpointsController interface {
	Create()
}

type privateEndpointsController struct {
	svcAdd endpoint.AddEndpointService
}

func NewPrivateEndpointsController(svcAdd endpoint.AddEndpointService) *privateEndpointsController {
	return &privateEndpointsController{svcAdd: svcAdd}
}

func (e *privateEndpointsController) Create(w http.ResponseWriter, r *http.Request) {
	var entity dto.CreateEndpointDto
	json.NewDecoder(r.Body).Decode(&entity)

	err := e.svcAdd.AddEndpoint(entity)

	if err != nil {
		webapi.SendError(err, w)
		return
	}
}
