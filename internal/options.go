package internal

// Options
// is strago server configs.
type Options struct {
	// Enable server or disable it.
	Enable bool
	// Server Port.
	Port int
	// Service Type which can be http or https.
	Type string
	// BalancingType selects the parameter to balance services.
	BalancingType int
}

// NewOptions
// returns one option instance.
func NewOptions() *Options {
	return &Options{}
}

// DefaultOptions
// returns a default config set of strago.
func DefaultOptions() *Options {
	return &Options{
		Enable:        false,
		Port:          9370,
		Type:          "http",
		BalancingType: RequestsCount,
	}
}

// ProductionOptions
// returns a set of configs for production of strago server.
func ProductionOptions() *Options {
	return &Options{
		Enable:        true,
		Port:          9370,
		Type:          "https",
		BalancingType: BusyTime,
	}
}
