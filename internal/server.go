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
		http:        http_client.New(),
		logger:      logger.NewLogger(),
		services:    createServices(services),
	}
}

func (s *server) Open(ip string) error {
	return s.changeStatusForAService(ip, true)
}

func (s *server) Close(ip string) error {
	return s.changeStatusForAService(ip, false)
}

func (s *server) Start() error {
	address := ":" + strconv.Itoa(s.port)

	gin.SetMode(gin.ReleaseMode)

	app := gin.Default()

	app.Use(s.handleRequests)

	if err := app.Run(address); err != nil {
		return fmt.Errorf("register server failed: %w", err)
	}

	return nil
}
