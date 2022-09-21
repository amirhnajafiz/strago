package main

import "github.com/amirhnajafiz/strago"

func main() {
	server := strago.NewServer(
		strago.WithServices(
			strago.ProductionConfigs(),
			"127.0.0.1:5050",
			"127.0.0.1:5051",
		),
	)

	_ = server.Close("127.0.0.1:5050")

	if err := server.Start(); err != nil {
		panic(err)
	}
}
