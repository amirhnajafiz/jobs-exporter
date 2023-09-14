package main

import (
	"github.com/amirhnajafiz/job-monitor/internal/config"
	"github.com/amirhnajafiz/job-monitor/internal/handler"
	"github.com/amirhnajafiz/job-monitor/internal/k8s"
	"github.com/amirhnajafiz/job-monitor/internal/kafka"
)

func main() {
	// load configs
	cfg := config.Load()

	// open kafka connection
	k, err := kafka.GetConnection(cfg.Kafka)
	if err != nil {
		panic(err)
	}

	// create k8s client
	c, err := k8s.GetClient(cfg.K8S)
	if err != nil {
		panic(err)
	}

	// create new handler
	h := handler.Handler{
		Interval:  cfg.Interval,
		KafkaConn: k,
		K8SClient: c,
	}

	// start monitor
	h.Monitor()
}
