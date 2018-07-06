package main

import (
	"github.com/rs/cors"
	"net/http"
)

func main() {
	server := NewServer()
	server.routes()
	http.ListenAndServe(":8000", cors.Default().Handler(server.router))
}
