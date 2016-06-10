FROM golang:1.6

MAINTAINER Akash

EXPOSE 8081 8081

ADD . /snabar_service

RUN cd /snabar_service/src/main

ENV GOPATH=/snabar_service/

CMD ["go","run", "src/main/route.go"]