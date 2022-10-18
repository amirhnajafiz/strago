package main

import "github.com/amirhnajafiz/strago"

func main() {
	// creating a new strago server
	server := strago.NewServer(strago.DefaultOptions())

	// set services
	server.WithServices("127.0.0.1:5050", "127.0.0.1:5051")

	// enable server
	server.Enable()

	// start server
	if err := server.Start(); err != nil {
		panic(err)
	}
}
