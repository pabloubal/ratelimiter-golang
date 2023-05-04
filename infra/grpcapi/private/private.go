package private

import (
	"context"
	"io/github/pabloubal/ratelimiter/dto"
	"io/github/pabloubal/ratelimiter/internal/usecases/endpoint"
	pb "io/github/pabloubal/ratelimiter/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedRateLimiterServer
	svcAdd endpoint.AddEndpointService
}

func (s *server) CreateEndpoint(ctx context.Context, req *pb.CreateEndpointRQ) (*pb.CreateEndpointRS, error) {
	dtoObj := dto.CreateEndpointDto{
		Path:   req.Path,
		Url:    req.Url,
		Limit:  int(req.Limit),
		Method: req.Method,
	}

	err := s.svcAdd.AddEndpoint(dtoObj)

	if err != nil {
		return nil, err
	}

	return &pb.CreateEndpointRS{Msg: "Endpoint created"}, nil
}

func Serve(svcAdd endpoint.AddEndpointService) *grpc.Server {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic("Cannot listen to 50051. " + err.Error())
	}

	srv := grpc.NewServer()
	pb.RegisterRateLimiterServer(srv, &server{svcAdd: svcAdd})

	go srv.Serve(listener)

	log.Println("Listening to :50051")

	return srv
}
