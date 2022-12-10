package config

import (
	"github.com/official-stallion/stallion-black-box-exporter/internal/client"
)

// Default configs.
func Default() Config {
	return Config{
		Client: client.Config{
			Address: "",
			Auth: client.Auth{
				Username: "root",
				Password: "",
			},
			Port: 8080,
		},
	}
}
