package main

import (
	"context"
	"io/github/pabloubal/ratelimiter/infra/webapi/privateapi"
	"io/github/pabloubal/ratelimiter/infra/webapi/publicapi"
	"io/github/pabloubal/ratelimiter/internal/domain"
	"io/github/pabloubal/ratelimiter/internal/usecases/endpoint"
	"log"
	"os"
)

func main() {
	ctx := context.Background()
	doneChan := make(chan os.Signal, 1)

	limiter := domain.NewLimiter()

	ucAddEndpoint := endpoint.NewAddEndpointUseCase(limiter)
	ucRQEndpoint := endpoint.NewRequestEndpointUseCase(limiter)

	srv := publicapi.Serve(ucRQEndpoint)
	adminSrv := privateapi.Serve(ucAddEndpoint)

	<-doneChan
	adminSrv.Shutdown(ctx)
	srv.Shutdown(ctx)
	log.Println("Done")
}
