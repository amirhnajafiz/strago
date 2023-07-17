package internal

import (
	"fmt"
	"strconv"

	"github.com/amirhnajafiz/strago/internal/metrics"
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
	// WithServices adds list of services to strago.
	WithServices(services ...string)
}

// NewServer
// creates a new load-balancer server.
func NewServer(opt *Options) LoadBalancer {
	server := newServer(
		opt.Port,
		opt.BalancingType,
		opt.Type,
	)

	return server
}

// server
// is the core of strago which manages the load-balancing.
type server struct {
	// type of strago server.
	serviceType string
	// server port.
	port int

	// balancing type.
	balancingType int

	// metrics of the server.
	metrics metrics.Metrics
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
	port int,
	balancingType int,
	serviceType string,
) *server {
	return &server{
		port:        port,
		serviceType: serviceType,

		balancingType: balancingType,

		metrics: metrics.NewMetrics(),
		http:    http_client.NewClient(),
		logger:  logger.NewLogger(),
	}
}

// WithServices
// adds services to strago server.
func (s *server) WithServices(services ...string) {
	s.services = generateServicesFromGiven(services)
}

// Start
// starting strago server.
func (s *server) Start() error {
	// change gin mode
	gin.SetMode(gin.ReleaseMode)

	address := ":" + strconv.Itoa(s.port)
	app := gin.Default()

	app.Use(s.handleRequests)
	app.GET("/metrics", s.metrics.Export)

	s.logger.Info("load balancer started", zap.String("port", address))

	if err := app.Run(address); err != nil {
		return fmt.Errorf("register server failed: %w", err)
	}

	return nil
}
