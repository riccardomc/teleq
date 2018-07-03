FROM golang:1.9-alpine AS build
RUN apk --no-cache add git make && \
  go get -u github.com/golang/dep/cmd/dep
WORKDIR $GOPATH/src/app
COPY . $GOPATH/src/app
RUN dep ensure && \
  make teleq-static

FROM scratch
COPY --from=build /go/src/app/teleq-static /teleq-static
ENTRYPOINT ["/teleq-static", "server"]