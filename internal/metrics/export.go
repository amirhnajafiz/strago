package metrics

import (
	"net/http"

	"github.com/amirhnajafiz/strago/pkg/model"

	"github.com/gin-gonic/gin"
)

// Export function sets metrics on http endpoint.
func (m *Metrics) Export(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, model.ExportMetrics{
		Requests:           m.numberOfRequests,
		FailedRequests:     m.numberOfFailedRequests,
		RequestsPerService: m.numberOfRequestsPerService,
		ResponseTime:       m.responseTime,
	})
}
