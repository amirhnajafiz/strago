package internal

import "github.com/amirhnajafiz/strago/pkg/model"

// Options
// is strago server configs.
type Options struct {
	// Port of http server.
	Port int
	// Secure Type which can be http or https.
	Secure bool
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
		Port:          9370,
		Secure:        false,
		BalancingType: model.RequestsCount,
	}
}

// ProductionOptions
// returns a set of configs for production of strago server.
func ProductionOptions() *Options {
	return &Options{
		Port:          9370,
		Secure:        true,
		BalancingType: model.BusyTime,
	}
}
