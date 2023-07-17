package internal

// Options
// is strago server configs.
type Options struct {
	Port          int
	Secure        bool
	BalancingType int
}

// NewOptions
// returns one option instance.
func NewOptions() *Options {
	return &Options{}
}
