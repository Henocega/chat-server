package config

import (
	"github.com/joho/godotenv"
)

// Load some
func Load(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}

	return nil
}

// GRPCConfig 1
type GRPCConfig interface {
	Address() string
}

// PGConfig 2
type PGConfig interface {
	DSN() string
}
