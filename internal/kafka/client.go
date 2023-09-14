package kafka

import (
	"context"

	"github.com/segmentio/kafka-go"
)

func GetConnection(cfg Config) (*kafka.Conn, error) {
	return kafka.DialLeader(context.Background(), "tcp", cfg.Host, cfg.Topic, cfg.Partition)
}
