package main

import (
	"net/http"
)

func main() {
	// The Server struct is essentially a server config
	// If the handler is nil, the default multiplexer DefaultServeMux is used
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: nil,
	}

	server.ListenAndServe()
}
