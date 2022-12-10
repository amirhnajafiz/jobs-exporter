package telemetry

// Config stores data for metrics server.
type Config struct {
	Address string `koanf:"address"`
	Enabled bool   `koanf:"enabled"`
}
