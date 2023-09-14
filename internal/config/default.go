package config

import (
	"github.com/amirhnajafiz/job-monitor/internal/k8s"
	"github.com/amirhnajafiz/job-monitor/internal/kafka"
)

func Default() Config {
	return Config{
		Interval: 5,
		Kafka: kafka.Config{
			Partition: 0,
			Topic:     "jm-jobs",
			Host:      "localhost:9292",
		},
		K8S: k8s.Config{
			Path:      "config",
			Namespace: "default",
		},
	}
}
