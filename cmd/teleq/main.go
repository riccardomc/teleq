package main

import (
	"flag"
	"net/http"
	"strconv"

	"github.com/riccardomc/teleq/server"
)

func main() {
	port := flag.Int("p", 9009, "port number")
	flag.Parse()

	server := server.NewStackServer()
	http.ListenAndServe(":"+strconv.Itoa(*port), server.Router)
}
