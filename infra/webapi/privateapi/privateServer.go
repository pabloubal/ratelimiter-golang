package privateapi

import (
	"io/github/pabloubal/ratelimiter/internal/usecases/endpoint"
	"log"
	"net/http"
)

func Serve(svcAdd endpoint.AddEndpointService) *http.Server {
	mux := createRoutes(svcAdd)

	srv := &http.Server{
		Addr:    ":8081",
		Handler: mux,
	}

	go srv.ListenAndServe()
	log.Println("Listening on :8080")

	return srv
}

func createRoutes(svcAdd endpoint.AddEndpointService) *http.ServeMux {
	mux := http.NewServeMux()
	endpointsController := NewPrivateEndpointsController(svcAdd)

	mux.HandleFunc("/endpoints", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			endpointsController.Create(w, r)
		default:
			handleNotAllowed(w, r)
			return
		}
	})

	return mux
}

func handleNotAllowed(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	return
}
