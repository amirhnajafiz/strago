package internal

import (
	"fmt"
	"net/http"
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
		opt.Debug,
	)

	return server
}

// server
// is the core of strago which manages the load-balancing.
type server struct {
	balancingType int
	port          int
	secure        bool
	debug         bool
	metrics       metrics.Metrics
	http          *client.HTTPClient
	logger        *zap.Logger
	services      []*service
}

// NewServer
// creates a new strago server.
func newServer(
	port,
	balancingType int,
	secure,
	debug bool,
) *server {
	level := zap.DebugLevel
	if !debug {
		level = zap.InfoLevel
	}

	return &server{
		port:          port,
		secure:        secure,
		balancingType: balancingType,
		debug:         debug,
		metrics:       metrics.NewMetrics(),
		http:          client.NewClient(),
		logger:        logger.NewLogger(level),
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
	if !s.debug {
		gin.SetMode(gin.ReleaseMode)
	}

	app := gin.Default()

	app.GET("/metrics", s.metrics.Export)
	app.GET("/health", func(context *gin.Context) {
		context.Status(http.StatusOK)
	})

	app.Use(s.handleRequests)

	s.logger.Info("load balancer started",
		zap.Int("port", s.port),
		zap.Bool("secure", s.secure),
		zap.Int("type", s.balancingType),
	)

	if err := app.Run(":" + strconv.Itoa(s.port)); err != nil {
		return fmt.Errorf("register server failed: %w", err)
	}

	return nil
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
