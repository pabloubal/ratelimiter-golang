package publicapi

import (
	"io/github/pabloubal/ratelimiter/internal/usecases/endpoint"
	"log"
	"net/http"
)

func Serve(svcRQ endpoint.RequestEndpointService) *http.Server {
	mux := createRoutes(svcRQ)
	srv := &http.Server{
		Addr:    ":80",
		Handler: mux,
	}

	go srv.ListenAndServe()
	log.Println("Listening on :80")

	return srv
}

func createRoutes(svcRQ endpoint.RequestEndpointService) *http.ServeMux {
	mux := http.NewServeMux()
	endpointsController := NewPublicEndpointsController(svcRQ)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		endpointsController.Request(w, r)
	})

	return mux
}
