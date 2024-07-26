package env

import (
	"errors"
	"os"
)

const (
	dsnEnvName = "PG_DSN"
)

type PGConfig struct {
	dsn string
}

func NewPGConfig() (*PGConfig, error) {
	dsn := os.Getenv(dsnEnvName)
	if len(dsn) == 0 {
		return nil, errors.New("environment variable `PG_DSN` is not set")
	}

	return &PGConfig{dsn: dsn}, nil
}

func (c *PGConfig) DSN() string { return c.dsn }
