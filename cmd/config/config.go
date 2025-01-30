package config

type Config struct {
	ServerAddress string
	BaseURL       string
}

func NewConfig() *Config {
	return &Config{}
}
