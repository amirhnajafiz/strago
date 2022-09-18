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
	// Open one of the services.
	Open(ip string) error
	// Close one of the services.
	Close(ip string) error
	// BanIP into blacklist.
	BanIP(ip string, version ...int) error
	// RecoverIP from blacklist.
	RecoverIP(ip string, version ...int) error
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
