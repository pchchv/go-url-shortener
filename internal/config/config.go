package config

import (
	"github.com/pchchv/env"
	"github.com/pchchv/golog"
)

func init() {
	// Load values from .env into the system
	if err := env.Load(); err != nil {
		golog.Panic("No .env file found")
	}
}
