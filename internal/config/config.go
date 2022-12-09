package config

import (
	"github.com/official-stallion/stallion-black-box-exporter/internal/client"
)

// Config stores data for black box configs
type Config struct {
	Client client.Config `koanf:"client"`
}

// Load configs.
func Load() Config {
	return Default()
}
