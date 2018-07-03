SOURCES=$(shell find . -name '*.go')
.PHONY: test deps deps-test run build install

build: deps teleq

teleq: $(SOURCES)
	go build -o $@ ./cmd/teleq/main.go ./cmd/teleq/commands.go

teleq-static:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o $@ \
				./cmd/teleq/main.go ./cmd/teleq/commands.go

install: teleq
	go install ./cmd/teleq/

deps:
	dep ensure

run: teleq
	./teleq

test: deps
	go test -v ./...
