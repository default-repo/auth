package main

import (
	"github.com/default-repo/auth/internal/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"

	desc "github.com/default-repo/auth/pkg/proto/auth_v1"
)

func main() {
	listener, err := net.Listen("tcp", server.GRPCPort)
	if err != nil {
		log.Fatalf("listener creating failed: %s", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterAuthV1Server(s, &server.GRPCServer{})

	log.Printf("grpc server started on %s", listener.Addr().String())

	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
