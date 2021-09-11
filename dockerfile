FROM golang:1.17-alpine
MAINTAINER Rytia <admin@zzfly.net>

WORKDIR /go/src/ip-service
COPY ./storage /go/src/ip-service/storage
COPY ./src /go/src/ip-service/src
COPY ./go.mod /go/src/ip-service/

RUN go env -w GO111MODULE="on"
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go mod download
RUN go get github.com/zzfly256/ip-service/src
RUN go build -o ./ip-service github.com/zzfly256/ip-service/src

EXPOSE 80/tcp
ENTRYPOINT ["/go/src/ip-service/ip-service"]