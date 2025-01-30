package config

type ConfigData struct {
	serverAddress string
	baseURL       string
}

func NewConfig() *ConfigData {
	return &ConfigData{}
}
