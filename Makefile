SOURCES=$(shell find . -name '*.go')
.PHONY: test deps deps-test run build install

build: deps teleq

teleq: $(SOURCES)
	go build -o $@ ./cmd/teleq/main.go ./cmd/teleq/commands.go

install: teleq
	go install ./cmd/teleq/

deps:
	dep ensure

run: teleq
	./teleq

test: deps
	go test -v ./...
