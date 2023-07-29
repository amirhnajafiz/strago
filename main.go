package main

import (
	"os"
	"strconv"
	"strings"

	"github.com/amirhnajafiz/strago/internal"
)

var (
	Services    string
	Port        int
	Secure      bool
	BalanceType int
	Debug       bool
)

func readEnv() {
	Services = os.Getenv("STRAGO_SERVICES")
	Port, _ = strconv.Atoi(os.Getenv("STRAGO_PORT"))
	BalanceType, _ = strconv.Atoi(os.Getenv("STRAGO_TYPE"))

	tmp := os.Getenv("STRAGO_SECURE")
	if tmp == "true" {
		Secure = true
	}

	tmp = os.Getenv("STRAGO_DEBUG")
	if tmp == "true" {
		Debug = true
	}
}

func main() {
	// get env variables
	readEnv()

	// list the services
	list := strings.Split(Services, "&")

	// create a new config
	config := internal.NewOptions()

	config.Secure = Secure
	config.Port = Port
	config.BalancingType = BalanceType
	config.Debug = Debug

	// create new server
	server := internal.NewServer(config)
	server.WithServices(list...)

	// start server
	if err := server.Start(); err != nil {
		panic(err)
	}
}
