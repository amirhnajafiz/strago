package internal

import (
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
)

func (s *server) getIP() string {
	ip := s.services[0]

	sort.Slice(s.services, func(i, j int) bool {
		return i < j
	})

	return ip
}

func (s *server) handleRequests(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, ctx.Request)
}
