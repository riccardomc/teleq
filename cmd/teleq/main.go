package main

import (
	"os"
)

func main() {
	app := New()
	app.Run(os.Args)
}
