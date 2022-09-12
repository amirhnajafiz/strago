package strago

type Config struct {
	Enable   bool
	Port     int
	Services []string
}

func WithDefaultConfigs() Config {
	return Config{
		Enable: false,
		Port:   9370,
	}
}
