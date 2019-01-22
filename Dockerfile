FROM golang:latest

RUN go get gopkg.in/redis.v4

WORKDIR /app
