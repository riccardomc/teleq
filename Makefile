test: deps
	go test -v ./...

deps: 
	go get "github.com/julienschmidt/httprouter"

teleq: server/server.go
	go build

run: teleq
	./teleq
