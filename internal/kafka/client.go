package kafka

import (
	"context"

	"github.com/segmentio/kafka-go"
)

func GetConnection(host, topic string, partition int) (*kafka.Conn, error) {
	return kafka.DialLeader(context.Background(), "tcp", host, topic, partition)
}
