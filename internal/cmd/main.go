package cmd

import (
	"time"

	"github.com/official-stallion/stallion-black-box-exporter/internal/client"
	"github.com/official-stallion/stallion-black-box-exporter/internal/config"
	"github.com/official-stallion/stallion-black-box-exporter/internal/telemetry/metrics"
)

const (
	topic = "stallion-black-box"
)

// Execute method of blackbox.
func Execute() {
	// load configs
	cfg := config.Load()
	// create a new channel
	channel := make(chan []byte)

	// enabling metrics server
	metrics.NewServer(cfg.Telemetry)

	// getting the metrics.
	blackBoxMetrics := metrics.NewMetrics()

	// activate client for consuming.
	{
		go func() {
			// creating a new client
			cli := client.Client{
				Cfg: cfg.Client,
			}

			// connecting client
			for {
				if er := cli.Connect(); er != nil {
					blackBoxMetrics.ConsumeErrors.Add(1)

					time.Sleep(2 * time.Second)

					continue
				}

				break
			}

			// start consuming over stallion
			cli.Subscribe(topic, channel)
		}()
	}

	// activate client for providing.
	{
		go func() {

		}()
	}
}
