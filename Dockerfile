FROM golang:latest

RUN go get gopkg.in/redis.v4 \
    && go get -u github.com/go-sql-driver/mysql

WORKDIR /app
