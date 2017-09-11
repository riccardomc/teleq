test: deps
	go test -v ./...

deps: 
	go get "github.com/julienschmidt/httprouter"
