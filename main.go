package main

import (
	"net/http"

	"github.com/riccardomc/teleq/server"
)

func main() {
	server := server.NewStackServer()
	http.ListenAndServe(":9009", server.Router)
}
