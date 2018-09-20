FROM golang:1.9-alpine AS build
RUN apk --no-cache add git make && \
  go get -u github.com/golang/dep/cmd/dep
RUN mkdir -p $GOPATH/src/github.com/riccardomc/teleq
WORKDIR $GOPATH/src/github.com/riccardomc/teleq
COPY . $GOPATH/src/github.com/riccardomc/teleq
RUN dep ensure && \
  make teleq-static

FROM scratch
COPY --from=build /go/src/github.com/riccardomc/teleq/teleq-static /teleq-static
ENTRYPOINT ["/teleq-static", "server"]
