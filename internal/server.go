package internal

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type server struct {
}

func (s *server) Start(port int) error {
	address := ":" + strconv.Itoa(port)
	app := gin.Default()

	gin.SetMode(gin.ReleaseMode)
	app.Use(handleRequests)

	if err := app.Run(address); err != nil {
		return fmt.Errorf("register server failed: %w", err)
	}

	return nil
}
