package strago

import (
	"github.com/amirhnajafiz/strago/internal"
)

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
	// Open one of the services.
	Open(ip string) error
	// Close one of the services.
	Close(ip string) error
}

// NewServer
// creates a new load-balancer server.
func NewServer(cfg *Config) LoadBalancer {
	server := internal.NewServer(
		cfg.Enable,
		cfg.Port,
		cfg.Type,
		cfg.Services...,
	)

	return server
}
