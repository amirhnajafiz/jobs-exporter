package client

import (
	"fmt"

	"github.com/official-stallion/stallion"
)

// Client
// manages the stallion golang skd connection
// for blackbox status reporting.
type Client struct {
	Cfg        Config
	Connection stallion.Client
}

// Connect
// to stallion server.
func (c *Client) Connect() error {
	conn, err := stallion.NewClient(c.Cfg.Host + ":" + c.Cfg.Port)
	if err != nil {
		return fmt.Errorf("stallion connection failed %w", err)
	}

	c.Connection = conn

	return nil
}

// Publish
// send message over stallion.
func (c *Client) Publish(topic string, payload []byte) error {
	return c.Connection.Publish(topic, payload)
}

// Subscribe
// subscribe over stallion.
func (c *Client) Subscribe(topic string, channel chan []byte) {
	c.Connection.Subscribe(topic, func(bytes []byte) {
		channel <- bytes
	})
}
