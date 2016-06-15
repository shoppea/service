FROM golang:1.6

MAINTAINER Akash

EXPOSE 8081 8081

ADD . /snabar_service

WORKDIR /snabar_service/src/main

ENV GOPATH=/snabar_service/

CMD ["go","run", "/snabar_service/src/main/route.go"]