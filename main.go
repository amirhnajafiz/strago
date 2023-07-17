package main

import (
	"flag"
	"strings"

	"github.com/amirhnajafiz/strago/internal"
)

func main() {
	var (
		services    = flag.String("services", "", "list of services")
		port        = flag.Int("port", 8080, "http port")
		secure      = flag.Bool("secure", false, "http secure or not")
		balanceType = flag.Int("type", 1, "load balancing type (1 request / 2 busy time)")
	)

	// parse flags
	flag.Parse()

	// list the services
	list := strings.Split(*services, "&")

	// create a new config
	config := internal.NewOptions()

	config.Secure = *secure
	config.Port = *port
	config.BalancingType = *balanceType

	// create new server
	server := internal.NewServer(config)
	server.WithServices(list...)

	// start server
	if err := server.Start(); err != nil {
		panic(err)
	}
}
