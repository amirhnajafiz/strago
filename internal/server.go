package internal

// LoadBalancer
// is the strago server which manages the
// features that a load-balancer has.
type LoadBalancer interface {
	// Start server.
	Start() error
	// Enable load balancing server.
	Enable()
	// Disable load balancing server.
	Disable()
	// WithServices adds list of services to strago.
	WithServices(services ...string)
	// Open one of the services.
	Open(ip string) error
	// Close one of the services.
	Close(ip string) error
}

// NewServer
// creates a new load-balancer server.
func NewServer(opt *Options) LoadBalancer {
	server := NewServer(
		opt.Enable,
		opt.Port,
		opt.BalancingType,
		opt.Type,
	)

	return server
}
