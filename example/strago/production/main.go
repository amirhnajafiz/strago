package main

import "github.com/amirhnajafiz/strago"

func main() {
	// creating a new server
	server := strago.NewServer(strago.ProductionOptions())

	// setting the services
	server.WithServices("127.0.0.1:5050", "127.0.0.1:5051")

	// closing one port
	_ = server.Close("127.0.0.1:5050")

	// starting server
	if err := server.Start(); err != nil {
		panic(err)
	}
}
