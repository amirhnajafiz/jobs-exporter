package client

// Config stores config information for clients.
type Config struct {
	Address string `koanf:"address"`
	Auth    Auth   `koanf:"auth"`
	Port    int    `koanf:"port"`
}

// Auth stores data needed for stallion authentication.
type Auth struct {
	Username string `koanf:"username"`
	Password string `koanf:"password"`
}
