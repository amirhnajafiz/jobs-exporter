package config

import (
	"log"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/structs"
	"github.com/official-stallion/stallion-black-box-exporter/internal/client"
	"github.com/official-stallion/stallion-black-box-exporter/internal/telemetry"
)

// Config stores data for black box configs
type Config struct {
	Client    client.Config    `koanf:"client"`
	Telemetry telemetry.Config `koanf:"telemetry"`
}

// Load configs.
func Load() Config {
	var instance Config

	k := koanf.New(".")

	// load default configuration from file
	if err := k.Load(structs.Provider(Default(), "koanf"), nil); err != nil {
		log.Fatalf("error loading default: %s", err)
	}

	// load configuration from file
	if err := k.Load(file.Provider("config.yaml"), yaml.Parser()); err != nil {
		log.Printf("error loading config.yml: %s", err)
	}

	if err := k.Unmarshal("", &instance); err != nil {
		log.Fatalf("error unmarshalling config: %s", err)
	}

	return instance
}
