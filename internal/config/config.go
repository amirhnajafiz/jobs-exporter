package config

import (
	"github.com/amirhnajafiz/job-monitor/internal/k8s"
	"github.com/amirhnajafiz/job-monitor/internal/kafka"
)

type Config struct {
	Interval int          `koanf:"interval"`
	Kafka    kafka.Config `koanf:"kafka"`
	K8S      k8s.Config   `koanf:"cluster"`
}

func Load() {

}
