package server

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/ptypes/empty"

	desc "github.com/default-repo/auth/pkg/proto/auth_v1"
)

// GRPCPort - this is the port on which the gRPC server is running
const GRPCPort = ":50051"

// GRPCServer - structure for implementing .proto interface
type GRPCServer struct {
	desc.UnimplementedAuthV1Server
}

// Create - creates a new object based on a CreateRequest and returns a CreateResponse
func (s *GRPCServer) Create(_ context.Context, r *desc.CreateRequest) (*desc.CreateResponse, error) {
	fmt.Printf("[ create ] request: %+v\n", r)

	return &desc.CreateResponse{
		UUID: "1d132bab-2e43-4fbc-bc46-3148076d07cc",
	}, nil
}

// Get - gets an object based on a GetRequest
func (s *GRPCServer) Get(_ context.Context, r *desc.GetRequest) (*desc.GetResponse, error) {
	fmt.Printf("[ get ] request: %+v\n", r)

	admin := desc.Role_Admin

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
