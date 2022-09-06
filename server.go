package strago

import (
	"log"

	"github.com/amirhnajafiz/strago/internal"
)

func NewServer() {
	server := internal.NewServer()

	log.Println(server.Start(8080))
}
