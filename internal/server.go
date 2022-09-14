package internal

import (
	"fmt"
	"strconv"

	"github.com/amirhnajafiz/strago/pkg/http_client"
	"github.com/amirhnajafiz/strago/pkg/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type server struct {
	enabled     bool
	http        http_client.HTTPClient
	logger      *zap.Logger
	port        int
	services    []*service
	serviceType string
}

func NewServer(enabled bool, port int, serviceType string, services ...string) *server {
	return &server{
		enabled:     enabled,
		logger:      logger.NewLogger(),
		port:        port,
		services:    createServices(services),
		serviceType: serviceType,
	}
}

func (s *server) toggle(ip string, status bool) error {
	for _, service := range s.services {
		if service.ip == ip {
			service.enable = status

			return nil
		}
	}

	return fmt.Errorf("service not found")
}

func (s *server) Open(ip string) error {
	return s.toggle(ip, true)
}

func (s *server) Close(ip string) error {
	return s.toggle(ip, false)
}

func (s *server) Start() error {
	address := ":" + strconv.Itoa(s.port)
	app := gin.Default()

	gin.SetMode(gin.ReleaseMode)
	app.Use(s.handleRequests)

	if err := app.Run(address); err != nil {
		return fmt.Errorf("register server failed: %w", err)
	}

	return nil
}
