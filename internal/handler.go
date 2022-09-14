package internal

import (
	"log"
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
)

func (s *server) getIP() string {
	serv := s.services[0]

	serv.used++

	sort.Slice(s.services, func(i, j int) bool {
		return s.services[i].used < s.services[j].used
	})

	return serv.ip
}

func (s *server) handleRequests(ctx *gin.Context) {
	// check service enable/disable status
	if !s.enabled {
		ctx.Status(http.StatusNotFound)
	}

	ip := s.getIP()

	req := ctx.Request
	uri := s.serviceType + "://" + ip + req.URL.Path

	log.Printf("got url: %s", uri)

	ctx.JSON(http.StatusOK, ctx.Request)
}
