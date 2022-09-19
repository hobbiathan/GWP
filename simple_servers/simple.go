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

	/*
		What the server struct looks like

		type Server struct {
			Addr			string
			Handler			Handler
			ReadTimeout		time.Duration
			WriteTimeout	time.Duration
			MaxHeaderBytes	int
			TLSConfig		*tls.Config
			TLSNextProto	map[string]func(*Server, *tls.Conn, Handler)
			ConnState		func(net.Conn, ConnState)
			ErrorLog		*log.Logger
		}

	*/

	server.ListenAndServe()
}
