package main

import (
	"net/http"
)

func main() {
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: nil,
	}

	// Requires Cert and Key arguments
	// ZeroSSL is pretty good
	server.ListenAndServeTLS("cert.pem", "key.pem")
}
