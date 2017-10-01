SOURCES=$(shell find . -name '*.go')
.PHONY=test deps run build install

build: deps teleq

teleq: $(SOURCES)
	go test -v ./...
	go build -o $@ ./cmd/teleq/main.go

install: teleq
	go install ./cmd/teleq/

deps: 
	go get "github.com/julienschmidt/httprouter"

run: teleq
	./teleq
