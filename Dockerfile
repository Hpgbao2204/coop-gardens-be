FROM golang:1.24.0-alpine3.7 AS builder

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go build -o main cmd/main.go

CMD ["cmd/coop-gardens-be/main"]