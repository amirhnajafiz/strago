package strago

import (
	"github.com/amirhnajafiz/strago/internal"
)

type LoadBalancer interface {
	Start() error
	Open(ip string) error
	Close(ip string) error
}

func NewServer(cfg *Config) LoadBalancer {
	server := internal.NewServer(cfg.Enable, cfg.Port, cfg.Type, cfg.Services...)

	return server
}
