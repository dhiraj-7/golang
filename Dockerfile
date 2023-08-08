FROM golang:1.16 AS base

RUN apt update

RUN apt upgrade -y

WORKDIR /go/src/app

COPY *.go .

RUN go mod init

RUN go build -o main .

EXPOSE 80

CMD ["/go/src/app/main"]