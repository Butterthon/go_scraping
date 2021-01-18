FROM golang:1.14.4-alpine

ENV GO111MODULE on

ADD . /go_scraping

WORKDIR /go_scraping

RUN apk update \
    && apk add --no-cache git \
    && go build -o go_scraping ./main.go
