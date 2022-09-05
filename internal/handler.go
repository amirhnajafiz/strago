package internal

import "net/http"

func handleRequests(writer http.ResponseWriter, request *http.Request) {
	_, _ = writer.Write([]byte(request.URL.String()))
}
