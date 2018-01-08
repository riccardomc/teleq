SOURCES=$(shell find . -name '*.go')
.PHONY: test deps deps-test run build install

build: deps teleq

teleq: $(SOURCES)
	go build -o $@ ./cmd/teleq/main.go ./cmd/teleq/commands.go

install: teleq
	go install ./cmd/teleq/

deps:
	go get "github.com/julienschmidt/httprouter"
	go get "github.com/urfave/cli"

run: teleq
	./teleq

deps-test:
	go get "github.com/jarcoal/httpmock"

test: deps-test
	go test -v ./...
