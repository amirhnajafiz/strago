package strago

// Config
// is strago server configs.
type Config struct {
	// Enable server or disable it.
	Enable bool
	// Server Port.
	Port int
	// Services which is a list of available services.
	Services []string
	// Service Type which can be http or https.
	Type string
}

// WithDefaultConfigs
// returns a default config set of strago.
func WithDefaultConfigs() *Config {
	return &Config{
		Enable: false,
		Port:   9370,
		Type:   "http",
	}
}

// ProductionConfigs
// returns a set of configs for production of strago server.
func ProductionConfigs() *Config {
	return &Config{
		Enable: true,
		Port:   9370,
		Type:   "https",
	}
}

// WithServices
// is used to attach services into strago config.
// If you don't pass config, it will use default configs.
func WithServices(config *Config, services ...string) *Config {
	if config == nil {
		config = WithDefaultConfigs()
	}

	config.Services = services

	return config
}
