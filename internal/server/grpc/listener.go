package grpc

import (
	"errors"
	"net"

	"github.com/default-repo/auth/internal/config/env"
)

type Listener struct {
	NetListener *net.Listener
}

func NewListener(cnf *env.GRPCConfig) (*Listener, error) {
	listener, err := net.Listen("tcp", cnf.Address())
	if err != nil {
		return nil, errors.New("listener creating failed; error: " + err.Error())
	}

	if listener == nil {
		return nil, errors.New("listener is nil")
	}

	return &Listener{NetListener: &listener}, err
}
