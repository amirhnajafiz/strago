package strago

type Config struct {
	Enable   bool
	Port     int
	Services []string
	Type     string
}

func WithDefaultConfigs() Config {
	return Config{
		Enable: false,
		Port:   9370,
		Type:   "http",
	}
}
