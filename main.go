package main

import (
	"context"
	privgrpc "io/github/pabloubal/ratelimiter/infra/grpcapi/private"
	pubgrpc "io/github/pabloubal/ratelimiter/infra/grpcapi/public"
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
	protoPrivSrv := privgrpc.Serve(ucAddEndpoint)
	protoPubSrv := pubgrpc.Serve(ucRQEndpoint)

	<-doneChan
	adminSrv.Shutdown(ctx)
	srv.Shutdown(ctx)
	protoPrivSrv.Stop()
	protoPubSrv.Stop()
	log.Println("Done")
}
