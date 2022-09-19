package main

import (
	"net/http"
)

// Handler functions take a ResponseWriter as a param, and a pointer to a request as a the second
func index(w http.ResponseWriter, r *http.Request) {
	threads, err := data.Threads(); if err == nil {
		_, err := session(w, r)

		public_tmpl_files := []string{
			"templates/layout.html",
			"templates/public.navbar.html",
			"templates/index.html"
		}

		private_tmpl_files := []string{
			"templates/layout.html",
			"templates/private.navbar.html",
			"templates/index.html",
		}

		var templates *template.Template
		if err != nil {
			templates = template.Must(template.ParseFiles(private_tmpl_files...))
		} else {
			templates = template.Must(template.ParseFiles(public_tmpl_files...))
		}
		templates.ExecuteTEmplate(w, "layout", threads)
	}
}

func main() {

	// Making a multiplexer
	// Basically the equivalent to a Rails router
	mux := http.NewServeMux()

	// Create handler (FileServer) to serve files
	files := http.FileServer(http.Dir("/public"))

	// /static/css/bootstrap.min.css -> /css/boostrap.min.css
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	// Redirect root to handler func
	// Note: Handlers and Handler Functions are not the same
	mux.HandleFunc("/", index)

	// Other stuff that'll be explained in the next chapter
	mux.HandleFunc("/err", err)

	// User sessions
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/logout", logout)
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/signup_account", signupAccount)
	mux.HandleFunc("/authenticate", authenticate)

	// Thread CRUD
	mux.HandleFunc("/thread/new", newThread)
	mux.HandleFunc("/thread/create", createThread)
	mux.HandleFunc("/thread/post", postThread)
	mux.HandleFunc("/thread/read", readThread)

	// Create server instance listening on :8080 utilizing earlier mux
	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}

	server.ListenAndServe()
}
