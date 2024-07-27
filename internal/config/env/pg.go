package env

import (
	"fmt"
	"os"
)

const dsnEnvName = "PG_DSN"

type PGConfig struct {
	dsn string
}

func NewPGConfig() (*PGConfig, error) {
	dsn := os.Getenv(dsnEnvName)
	if len(dsn) == 0 {
		return nil, fmt.Errorf("environment variable %s is not set", dsnEnvName)
	}

	return &PGConfig{dsn: dsn}, nil
}

func (c *PGConfig) DSN() string { return c.dsn }
