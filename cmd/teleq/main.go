package main

import (
	"os"

	"github.com/riccardomc/teleq/server"
)

//Server used by teleq
var Server stackserver.ServerInterface

func main() {
	Server = stackserver.New()
	app := New()
	app.Run(os.Args)
}
