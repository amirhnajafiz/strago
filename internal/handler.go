package internal

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// handleRequests
// gets user inputs and processes them.
func (s *server) handleRequests(ctx *gin.Context) {
	// check service enable/disable status
	if !s.enabled {
		ctx.Status(http.StatusNotFound)
	}

	// load-balancing logic
	ip := s.getOneIPFromServices()

	// extract request and create a new address
	req := ctx.Request
	uri := s.serviceType + "://" + ip + req.URL.Path

	s.logger.Info("load balancer given ip", zap.String("uri", uri))

	// handle the request with new uri
	res, err := s.handle(uri, req)
	if err != nil {
		ctx.Status(res.StatusCode)

		_ = ctx.Error(err)

		return
	}

	ctx.JSON(res.StatusCode, res.Body)
}

// handle
// manages the user request, based on the method.
func (s *server) handle(uri string, req *http.Request) (*http.Response, error) {
	switch req.Method {
	case http.MethodGet:
		return s.http.Get(uri)
	case http.MethodPost:
		return s.http.Post(uri, req.Body)
	case http.MethodPut:
		return s.http.Put(uri, req.Body)
	case http.MethodDelete:
		return s.http.Delete(uri)
	default:
		return nil, fmt.Errorf("unsupported protocol")
	}
}
