package main

import (
    "net"
    "log"
    "context"
    "my_app/api"
    "google.golang.org/grpc/reflection"
    "google.golang.org/grpc"
)

type server struct{}

func (s *server)SayHello(ctx context.Context, args *api.HelloRequest) (*api.HelloResponse, error) {
    log.Printf("Name=%s\n", args.Name)
    return &api.HelloResponse{Data:"OK"}, nil
}

func main() {
    port:=":8080"
    lis, err := net.Listen("tcp", port)
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    log.Printf("Hello server listening at localhost%s", port)
    s := grpc.NewServer()
    api.RegisterHelloServer(s, &server{}) 
    reflection.Register(s)
        if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}

