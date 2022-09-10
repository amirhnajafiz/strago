package strago

type Config struct {
	Enable   bool
	Services []string
}

func WithDefaultConfigs() Config {
	return Config{
		Enable: false,
	}
}
