package strago

type Config struct {
	Enable   bool
	Port     int
	Services []string
	Type     string
}

func ProductionConfigs() *Config {
	return &Config{
		Enable: true,
		Port:   9370,
		Type:   "https",
	}
}

func WithDefaultConfigs() *Config {
	return &Config{
		Enable: false,
		Port:   9370,
		Type:   "http",
	}
}

func WithServices(config *Config, services ...string) *Config {
	if config == nil {
		cfg := WithDefaultConfigs()

		cfg.Services = services

		return cfg
	}

	config.Services = services

	return config
}
