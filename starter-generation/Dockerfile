FROM golang:1.19.2-alpine3.16

RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct

WORKDIR /go/server
COPY . .

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories
RUN apk update
RUN apk add gcc g++ libffi-dev make zlib-dev libcec-dev libtool

RUN go install github.com/swaggo/swag/cmd/swag

RUN swag init --pd

RUN go build -v -o server .

FROM alpine:3.16.2

WORKDIR /go/server

COPY --from=0 /go/server/server ./

ENTRYPOINT ./server
