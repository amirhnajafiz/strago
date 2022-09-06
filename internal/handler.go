package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleRequests(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, ctx.Request)
}
