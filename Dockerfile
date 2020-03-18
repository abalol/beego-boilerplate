
FROM golang:1.14-alpine

WORKDIR /app

ENV GOPATH $GOPATH:/go
ENV PATH $PATH:$GOPATH/bin

RUN apk add --update --no-cache git
RUN go get "github.com/beego/bee"

COPY . .

RUN go mod download
