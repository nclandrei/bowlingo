package main

import "net/http"

func main() {
	server := NewServer()
	server.routes()
	http.ListenAndServe(":8080", server.router)
}
