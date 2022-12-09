package client

// Config stores config information for clients.
type Config struct {
	Address  string `koanf:"address"`
	Consumer bool   `koanf:"consumer"`
}
