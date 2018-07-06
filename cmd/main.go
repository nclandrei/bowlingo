package main

import (
	"github.com/nclandrei/bowlingo"
	"github.com/rs/cors"
	"net/http"
)

func main() {
	server := bowlingo.NewServer()
	server.Routes()
	http.ListenAndServe(":8000", cors.Default().Handler(server))
}
