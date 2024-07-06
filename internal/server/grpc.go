package server

import (
	"context"
	"fmt"
	desc "github.com/default-repo/auth/pkg/proto/auth_v1"
	"github.com/golang/protobuf/ptypes/empty"
)

const GRPCPort = ":50051"

type GRPCServer struct {
	desc.UnimplementedAuthV1Server
}

func (s *GRPCServer) Create(_ context.Context, r *desc.CreateRequest) (*desc.CreateResponse, error) {
	fmt.Printf("[ create ] request: %+v\n", r)

	return &desc.CreateResponse{
		UUID: "1d132bab-2e43-4fbc-bc46-3148076d07cc",
	}, nil
}

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

func (s *GRPCServer) Update(_ context.Context, r *desc.UpdateRequest) (*empty.Empty, error) {
	fmt.Printf("[ update ] request: %+v\n", r)

	return nil, nil
}

func (s *GRPCServer) Delete(_ context.Context, r *desc.DeleteRequest) (*empty.Empty, error) {
	fmt.Printf("[ delete ] request: %+v\n", r)

	return nil, nil
}
