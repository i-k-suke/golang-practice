FROM golang:1.11
LABEL name=golang-app

RUN mkdir /go/src/app

WORKDIR /go/src/app

RUN apt-get update
RUN apt-get install -y vim
RUN apt-get install -y less

ENV CGO_ENABLE 0
