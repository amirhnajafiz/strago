package internal

import (
	"fmt"
	"strconv"

	"github.com/amirhnajafiz/strago/pkg/http_client"
	"github.com/amirhnajafiz/strago/pkg/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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
	server := newServer(
		opt.Enable,
		opt.Port,
		opt.BalancingType,
		opt.Type,
	)

	return server
}

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
func newServer(
	enabled bool,
	port int,
	balancingType int,
	serviceType string,
) *server {
	return &server{
		enabled:     enabled,
		port:        port,
		serviceType: serviceType,

		balancingType: balancingType,

		metrics: newMetrics(),
		http:    http_client.NewClient(),
		logger:  logger.NewLogger(),
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

// WithServices
// adds services to strago server.
func (s *server) WithServices(services ...string) {
	s.services = generateServicesFromGiven(services)
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
