package internal

import "net/http"

type server struct {
}

func (s *server) register() {
	http.HandleFunc("/*", handleRequests)
}
