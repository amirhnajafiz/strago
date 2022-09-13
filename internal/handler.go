package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *server) handleRequests(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, ctx.Request)
}
