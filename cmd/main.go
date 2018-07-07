package main

/*
	The command for launching the server is stored in the cmd folder as it is considered
	best practice in the Go community - it differentiates the command from the library which
	can be used as a standalone product as well. For more details, please see here:
	https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1 and
	https://medium.com/@benbjohnson/structuring-applications-in-go-3b04be4ff091
*/

import (
	"flag"
	"fmt"
	"github.com/nclandrei/bowlingo"
	"github.com/rs/cors"
	"log"
	"net/http"
)

var (
	address = flag.String("addr", "0.0.0.0", "hostname for the bowling server")
	port    = flag.Int("port", 8000, "port for the bowling server")
)

func main() {
	flag.Parse()
	server := bowlingo.NewServer()
	server.Routes()
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", *address, *port), cors.Default().Handler(server)))
}
