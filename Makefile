SOURCES=$(shell find . -name '*.go')

.PHONY: build
build: deps teleq

teleq: $(SOURCES)
	go build -o $@ ./cmd/teleq/main.go ./cmd/teleq/commands.go

teleq-static:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o $@ \
				./cmd/teleq/main.go ./cmd/teleq/commands.go

.PHONY: install
install: teleq
	go install ./cmd/teleq/

.PHONY: deps
deps:
	dep ensure

.PHONY: run
run: teleq
	./teleq

.PHONY: test
test: deps
	go test -v ./...
