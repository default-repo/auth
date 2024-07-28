package grpc

import (
	"context"
	"fmt"

	desc "github.com/default-repo/auth/pkg/proto/auth_v1"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// GRPCServer - structure for implementing .proto interface
type GRPCServer struct {
	desc.UnimplementedAuthV1Server
}

type Server struct {
	S *grpc.Server
}

func NewGRPCServer() (*Server, error) {
	s := grpc.NewServer()

	reflection.Register(s)
	desc.RegisterAuthV1Server(s, &GRPCServer{})

	return &Server{
		S: s,
	}, nil
}

// Create - creates a new object based on a CreateRequest and returns a CreateResponse
func (s *GRPCServer) Create(_ context.Context, r *desc.CreateRequest) (*desc.CreateResponse, error) {
	fmt.Printf("[ create ] request: %+v\n", r)

	return &desc.CreateResponse{
		Uuid: "1d132bab-2e43-4fbc-bc46-3148076d07cc",
	}, nil
}

// Get - gets an object based on a GetRequest
func (s *GRPCServer) Get(_ context.Context, r *desc.GetRequest) (*desc.GetResponse, error) {
	fmt.Printf("[ get ] request: %+v\n", r)

	admin := desc.Role_ADMIN

	return &desc.GetResponse{
		User: &desc.User{
			Uuid:  "1d132bab-2e43-4fbc-bc46-3148076d07cc",
			Name:  "John Doe",
			Email: "johndoe@gmail.com",
			Role:  &admin,
		},
	}, nil
}

// Update - updates an element based on a UpdateRequest and returns an empty response
func (s *GRPCServer) Update(_ context.Context, r *desc.UpdateRequest) (*empty.Empty, error) {
	fmt.Printf("[ update ] request: %+v\n", r)

	return nil, nil
}

// Delete - deletes an element based on a DeleteRequest and returns an empty response
func (s *GRPCServer) Delete(_ context.Context, r *desc.DeleteRequest) (*empty.Empty, error) {
	fmt.Printf("[ delete ] request: %+v\n", r)

	return nil, nil
}
