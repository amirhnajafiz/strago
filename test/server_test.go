package test

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/amirhnajafiz/strago"
	"github.com/gin-gonic/gin"
)

func startTestServer(port string) {
	go func() {
		app := gin.Default()

		app.GET("/", func(context *gin.Context) {
			context.Status(http.StatusOK)
		})

		_ = app.Run(port)
	}()
}

func TestServer(t *testing.T) {
	opt := strago.NewOptions()

	opt.Enable = true
	opt.Port = 9370
	opt.Type = "http"
	opt.BalancingType = strago.RequestsCount

	// creating a new server
	server := strago.NewServer(opt)

	if server == nil {
		t.Error(fmt.Errorf("server initialize failed"))

		os.Exit(1)
	}

	// starting two servers
	startTestServer(":5050")
	startTestServer(":5051")

	// setting the services
	server.WithServices("127.0.0.1:5050", "127.0.0.1:5051")

	// starting strago server
	go func() {
		// starting server
		if err := server.Start(); err != nil {
			t.Error(err)

			os.Exit(1)
		}
	}()

	// fetch strago
	resp, err := http.Get("http://localhost:9370")
	if err != nil {
		t.Error(err)

		os.Exit(1)
	}

	// check result
	if resp.StatusCode != http.StatusOK {
		t.Error("failed to get instance")
	}
}
