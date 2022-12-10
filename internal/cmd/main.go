package cmd

import (
	"encoding/json"
	"time"

	"github.com/official-stallion/stallion-black-box-exporter/internal/client"
	"github.com/official-stallion/stallion-black-box-exporter/internal/config"
	"github.com/official-stallion/stallion-black-box-exporter/internal/telemetry/metrics"
)

const (
	topic = "stallion-black-box"
)

// Payload type stores data that we send over stallion.
type Payload struct {
	Ok      bool      `json:"ok"`
	Created time.Time `json:"created"`
}

// Execute method of blackbox.
func Execute() {
	// load configs
	cfg := config.Load()
	// create a new channel
	channel := make(chan []byte)

	// enabling metrics server
	metrics.NewServer(cfg.Telemetry).Start()

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

			for {
				// generate a payload
				p := Payload{
					Ok:      true,
					Created: time.Now(),
				}

				// convert to bytes
				bytes, _ := json.Marshal(p)

				// publish over stallion
				if er := cli.Publish(topic, bytes); er != nil {
					blackBoxMetrics.PublishErrors.Add(1)
				}

				// wait for 2 seconds
				time.Sleep(2 * time.Second)
			}
		}()
	}

	// handling read payloads
	for {
		select {
		case bytes := <-channel:
			// creating a payload
			var payload Payload
			// unmarshalling payload
			if err := json.Unmarshal(bytes, &payload); err != nil {
				blackBoxMetrics.ConsumeErrors.Add(1)
			}
			// check payload validation
			if payload.Ok {
				diff := time.Now().Sub(payload.Created)

				blackBoxMetrics.ResponseTime.Observe(diff.Seconds())
			}
		}
	}
}
