package strago

import (
	"github.com/amirhnajafiz/strago/internal"
)

type LoadBalancer interface {
	Start(int) error
}

func NewServer() LoadBalancer {
	server := internal.NewServer()

	return server
}
