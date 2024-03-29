package internal

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// handleRequests
// gets user inputs and processes them.
func (s *server) handleRequests(ctx *gin.Context) {
	s.metrics.IncRequest()

	// load-balancing logic
	selectedService := s.getOneIPFromServices()
	if selectedService == nil {
		s.logger.Debug("all services are disabled")

		_ = ctx.Error(fmt.Errorf("services are closed at the moment"))

		return
	}

	// extract request and create a new address
	req := ctx.Request

	prefix := "http"
	if s.secure {
		prefix = "https"
	}

	uri := prefix + "://" + selectedService.ip + req.URL.Path

	s.logger.Info("load balancer given ip", zap.String("uri", uri))
	s.metrics.IncRequestPer(selectedService.ip)

	// starting time
	start := time.Now()

	// handle the request with new uri
	res, err := s.handle(uri, req)
	if err != nil {
		s.logger.Debug("handle request failed", zap.Error(err))
		s.metrics.IncFailed()

		_ = ctx.Error(err)

		return
	}

	// creating a buffer for body
	buffer := make([]byte, 2048)
	_, _ = res.Body.Read(buffer)

	// calculating the response time
	duTime := time.Since(start)
	// accumulate the busy times of a service
	selectedService.busy = selectedService.busy + duTime

	s.logger.Info("response time", zap.Duration("duration", duTime))
	s.metrics.AddResponse(duTime.Minutes())

	// sending the service response
	ctx.Status(res.StatusCode)
	_, _ = ctx.Writer.Write(buffer)
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
