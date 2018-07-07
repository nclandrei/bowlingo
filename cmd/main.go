package main

import (
	"flag"
	"fmt"
	"github.com/nclandrei/bowlingo"
	"github.com/rs/cors"
	"net/http"
)

var (
	address = flag.String("addr", "localhost", "hostname for the bowling server")
	port    = flag.Int("port", 8000, "port for the bowling server")
)

func main() {
	flag.Parse()
	server := bowlingo.NewServer()
	server.Routes()
	http.ListenAndServe(fmt.Sprintf("%s:%d", *address, *port), cors.Default().Handler(server))
}
