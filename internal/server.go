package internal

import (
	"fmt"
	"strconv"

	"github.com/amirhnajafiz/strago/internal/metrics"
	"github.com/amirhnajafiz/strago/pkg/client"
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
		opt.Secure,
	)

	return server
}

// server
// is the core of strago which manages the load-balancing.
type server struct {
	balancingType int
	port          int
	secure        bool

	metrics  metrics.Metrics
	http     *client.HTTPClient
	logger   *zap.Logger
	services []*service
}

// NewServer
// creates a new strago server.
func newServer(
	port int,
	balancingType int,
	secure bool,
) *server {
	return &server{
		port:          port,
		secure:        secure,
		balancingType: balancingType,
		metrics:       metrics.NewMetrics(),
		http:          client.NewClient(),
		logger:        logger.NewLogger(),
	}
}

// generateServicesFromGiven
// creates the list of the services.
func generateServicesFromGiven(services []string) []*service {
	list := make([]*service, len(services))

	for index, ip := range services {
		list[index] = &service{
			ip:   ip,
			used: 0,
			busy: 0,
		}
	}

	return list
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
