// +build !appengine

// A stand-alone HTTP server providing a web UI for task management.
package main

import (
	"net/http"

	"github.com/campoy/todo/server"
)

func main() {
	server.RegisterHandlers()
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.ListenAndServe(":8080", nil)
}
