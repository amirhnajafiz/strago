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

	if err := server.BanIP("127.*.*.*"); err != nil {
		panic(err)
	}

	if err := server.Start(); err != nil {
		panic(err)
	}
}
