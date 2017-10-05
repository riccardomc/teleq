package main

import (
	"flag"

	"github.com/riccardomc/teleq/stackserver"
)

func main() {
	port := flag.Int("p", 9009, "port number")
	flag.Parse()

	config := &stackserver.ServerConfig{*port}

	s := stackserver.New(config)
	s.Serve()
}
