package main

import (
	"net/http"
)

func main() {

	// Making a multiplexer
	mux := http.NewServeMux()

	// Create handler (FileServer) to serve files
	files := http.FileServer(http.Dir("/public"))

	// /static/css/bootstrap.min.css -> /css/boostrap.min.css
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	// Redirect root to handler func
	// Note: Handlers and Handler Functions are not the same
	mux.HandleFunc("/", index)

	// Create server instance listening on :8080 utilizing earlier mux
	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}

	server.ListenAndServe()
}
