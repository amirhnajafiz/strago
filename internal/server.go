package internal

import (
	"fmt"
	"strconv"

	"github.com/amirhnajafiz/strago/pkg/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type server struct {
	enabled     bool
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
