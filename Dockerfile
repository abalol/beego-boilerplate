
FROM golang:1.14.0

ENV GOPATH $GOPATH:/go
ENV PATH $PATH:$GOPATH/bin

ADD ./api /go/src/app/api

WORKDIR /go/src/app/api

RUN apt-get update
RUN go get "github.com/go-sql-driver/mysql"
RUN go get "github.com/beego/bee"
RUN go get "github.com/astaxie/beego"

RUN go build