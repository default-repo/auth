package env

import (
	"errors"
	"net"
	"os"
)

const (
	grpcHostEnvName = "GRPC_HOST"
	grpcPortEnvName = "GRPC_PORT"
)

type GRPCConfig struct {
	host string
	port string
}

func NewGRPCConfig() (*GRPCConfig, error) {
	host := os.Getenv(grpcHostEnvName)
	if len(host) == 0 {
		return nil, errors.New("environment variable `GRPC_HOST` is not set")
	}

	port := os.Getenv(grpcPortEnvName)
	if len(port) == 0 {
		return nil, errors.New("environment variable `GRPC_PORT` is not set")
	}

	return &GRPCConfig{
		host: host,
		port: port,
	}, nil
}

func (cfg *GRPCConfig) Address() string {
	return net.JoinHostPort(cfg.host, cfg.port)
}
