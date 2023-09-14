package kafka

type Config struct {
	Host      string `koanf:"host"`
	Topic     string `koanf:"topic"`
	Partition int    `koanf:"partition"`
}
