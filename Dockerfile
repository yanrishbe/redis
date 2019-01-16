FROM golang:1.11.1-alpine

WORKDIR /go/src/redis/main

COPY *.go /go/src/redis/main

RUN apk update \
&& apk upgrade \
&& apk add --no-cache git vim

RUN go get -u github.com/golang/lint/golint \
&& go get -u golang.org/x/tools/cmd/goimports

RUN mv /go/bin/* /usr/local/go/bin/

RUN go build *.go \
&& rm -rf *.go
