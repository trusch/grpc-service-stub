package greeter

import (
	"context"

	"github.com/trusch/grpc-service-stub/pkg/protobuf/greeter"
)

func New() (greeter.GreeterServer, error) {
	return &greeterServer{}, nil
}

type greeterServer struct {
	greeter.UnimplementedGreeterServer
}

func (g *greeterServer) Greet(ctx context.Context, req *greeter.GreetRequest) (*greeter.GreetResponse, error) {
	return &greeter.GreetResponse{
		Greeting: "Hello " + req.GetName(),
	}, nil
}
