package config

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/amirhnajafiz/job-monitor/internal/k8s"
	"github.com/amirhnajafiz/job-monitor/internal/kafka"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/structs"
	"github.com/tidwall/pretty"
)

const (
	// Prefix indicates environment variables prefix.
	Prefix = "JM_"
)

type Config struct {
	Interval int          `koanf:"interval"`
	Kafka    kafka.Config `koanf:"kafka"`
	K8S      k8s.Config   `koanf:"cluster"`
}

func Load() Config {
	var instance Config

	k := koanf.New(".")

	// load default configuration from file
	if err := k.Load(structs.Provider(Default(), "koanf"), nil); err != nil {
		log.Fatalf("error loading default: %s", err)
	}

	// load environment variables
	if err := k.Load(env.Provider(Prefix, ".", func(s string) string {
		return strings.ReplaceAll(strings.ToLower(
			strings.TrimPrefix(s, Prefix)), "__", ".")
	}), nil); err != nil {
		log.Printf("error loading environment variables: %s", err)
	}

	if err := k.Unmarshal("", &instance); err != nil {
		log.Fatalf("error unmarshalling config: %s", err)
	}

	indent, err := json.MarshalIndent(instance, "", "\t")
	if err != nil {
		log.Fatalf("error marshaling config to json: %s", err)
	}

	indent = pretty.Color(indent, nil)
	tmpl := `
	================ Loaded Configuration ================
	%s
	======================================================
	`
	log.Printf(tmpl, string(indent))

	return instance
}
