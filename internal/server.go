package internal

import (
	"fmt"
	"strconv"

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

	// balancing type.
	balancingType int

	// metrics of the server.
	metrics metrics
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
	balancingType int,
	serviceType string,
	services ...string,
) *server {
	return &server{
		enabled:     enabled,
		port:        port,
		serviceType: serviceType,

		balancingType: balancingType,

		metrics:  newMetrics(),
		http:     http_client.NewClient(),
		logger:   logger.NewLogger(),
		services: generateServicesFromGiven(services),
	}
}

// Enable
// allow load balancing server to pass requests.
func (s *server) Enable() {
	s.enabled = true
}

// Disable
// block all the requests that are sent to server.
func (s *server) Disable() {
	s.enabled = false
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

// Start
// starting strago server.
func (s *server) Start() error {
	// change gin mode
	gin.SetMode(gin.ReleaseMode)

	address := ":" + strconv.Itoa(s.port)
	app := gin.Default()

	app.Use(s.handleRequests)
	app.GET("/metrics", s.prometheusHandler())

	s.logger.Info("load balancer started", zap.String("port", address))

	if err := app.Run(address); err != nil {
		return fmt.Errorf("register server failed: %w", err)
	}

	return nil
}
