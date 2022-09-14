package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (s *server) handleRequests(ctx *gin.Context) {
	// check service enable/disable status
	if !s.enabled {
		ctx.Status(http.StatusNotFound)
	}

	ip := s.getIP()

	req := ctx.Request
	uri := s.serviceType + "://" + ip + req.URL.Path

	s.logger.Info("new url", zap.String("uri", uri))

	ctx.JSON(http.StatusOK, ctx.Request)
}
