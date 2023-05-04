package public

import (
	"context"
	"io/github/pabloubal/ratelimiter/internal/mapper"
	"io/github/pabloubal/ratelimiter/internal/usecases/endpoint"
	pb "io/github/pabloubal/ratelimiter/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedRateLimiterServer
	svcRQ endpoint.RequestEndpointService
}

func (s *server) RequestEndpoint(ctx context.Context, req *pb.RequestEndpointRQ) (*pb.RequestEndpointRS, error) {
	entity, mapErr := mapper.GrpcRQToEntity(req)
	if mapErr != nil {
		return nil, mapErr
	}

	svcResponse, err := s.svcRQ.RequestEndpoint(entity)
	if err != nil {
		return nil, err
	}

	return mapper.HttpResponseToGrpc(svcResponse), nil
}

func Serve(svcRQ endpoint.RequestEndpointService) *grpc.Server {
	listener, err := net.Listen("tcp", ":50050")
	if err != nil {
		panic("Cannot listen to 50050. " + err.Error())
	}

	srv := grpc.NewServer()
	pb.RegisterRateLimiterServer(srv, &server{svcRQ: svcRQ})

	go srv.Serve(listener)

	log.Println("Listening to :50050")
	return srv
}
