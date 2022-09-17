package internal

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/amirhnajafiz/strago/pkg/http_client"
	"github.com/amirhnajafiz/strago/pkg/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// server
// is the core of strago which manages the load-balancing.
type server struct {
	// enable/disable status.
	enabled bool
	// type of strago server.
	serviceType string
	// server port.
	port int

	// blacklist ips for firewall
	blacklist []string

	// http client instance.
	http *http_client.HTTPClient
	// logger instance.
	logger *zap.Logger
	// list of the services.
	services []*service
}

// NewServer
// creates a new strago server.
func NewServer(
	enabled bool,
	port int,
	serviceType string,
	services ...string,
) *server {
	return &server{
		enabled:     enabled,
		port:        port,
		serviceType: serviceType,

		http:     http_client.NewClient(),
		logger:   logger.NewLogger(),
		services: generateServicesFromGiven(services),
	}
}

// Open
// allow strago to send requests to a service.
func (s *server) Open(ip string) error {
	return s.changeStatusForAService(ip, true)
}

// Close
// disallow strago to send requests to a service.
func (s *server) Close(ip string) error {
	return s.changeStatusForAService(ip, false)
}

// BanIP
// adds an IP into blacklist of our service.
func (s *server) BanIP(ip string) error {
	for _, part := range strings.Split(ip, ".") {
		if _, err := strconv.Atoi(part); err != nil && part != "*" {
			return fmt.Errorf("wrong ip format")
		}
	}

	s.blacklist = append(s.blacklist, ip)

	return nil
}

// RecoverIP
// removes an IP from server blacklist.
func (s *server) RecoverIP(ip string) error {
	for index, blackListIP := range s.blacklist {
		if blackListIP == ip {
			s.blacklist = append(s.blacklist[:index], s.blacklist[index+1:]...)

			return nil
		}
	}

	return fmt.Errorf("ip not found")
}

// Start
// starting strago server.
func (s *server) Start() error {
	// change gin mode
	gin.SetMode(gin.ReleaseMode)

	address := ":" + strconv.Itoa(s.port)
	app := gin.Default()

	v1 := app.Use(s.firewallHandler)
	v1.Use(s.handleRequests)

	if err := app.Run(address); err != nil {
		return fmt.Errorf("register server failed: %w", err)
	}

	return nil
}
