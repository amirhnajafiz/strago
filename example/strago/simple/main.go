package main

import "github.com/amirhnajafiz/strago"

func main() {
	server := strago.NewServer(
		strago.WithServices(
			strago.WithDefaultConfigs(),
			"127.0.0.1:5050",
			"127.0.0.1:5051",
		),
	)

	server.Enable()

	if err := server.Start(); err != nil {
		panic(err)
	}
}
