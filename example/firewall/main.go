package main

import "github.com/amirhnajafiz/strago"

func main() {
	// creating a new strago sever
	server := strago.NewServer(
		strago.WithServices(
			strago.WithDefaultConfigs(),
			// given services ips
			"127.0.0.1:5050",
			"127.0.0.1:5051",
		),
	)

	// ban any ip in range 127... version 4
	if err := server.BanIP("127.*.*.*"); err != nil {
		panic(err)
	}

	// ban any ip in range 12...12 version 6
	if err := server.BanIP("12:*:*:*:22:*:*:*"); err != nil {
		panic(err)
	}

	if err := server.Start(); err != nil {
		panic(err)
	}
}
