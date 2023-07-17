package internal

// Options
// is strago server configs.
type Options struct {
	Port          int
	BalancingType int
	Secure        bool
	Debug         bool
}

// NewOptions
// returns one option instance.
func NewOptions() *Options {
	return &Options{}
}
