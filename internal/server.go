package internal

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type server struct {
	enabled  bool
	port     int
	services []string
}

func NewServer(enabled bool, port int, services []string) *server {
	return &server{
		enabled:  enabled,
		port:     port,
		services: services,
	}
}

func (s *server) Start() error {
	address := ":" + strconv.Itoa(s.port)
	app := gin.Default()

	gin.SetMode(gin.ReleaseMode)
	app.Use(handleRequests)

	if err := app.Run(address); err != nil {
		return fmt.Errorf("register server failed: %w", err)
	}

	return nil
}
