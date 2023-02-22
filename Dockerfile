FROM golang:1.19.2-alpine3.16
ARG pgName

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

COPY . /go/clover

WORKDIR /go/clover/${pgName}

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories
RUN apk update
RUN apk add gcc g++ libffi-dev make zlib-dev libcec-dev libtool
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init --pd
RUN go build -v -o server .

FROM alpine:3.16.2
ARG pgName
ARG httpPort

WORKDIR /go/clover

ENV TZ Asia/Shanghai

COPY --from=0 /go/clover/${pgName}/server .
COPY --from=0 /go/clover/deploy/docker/clover/config/${pgName} ./configs

EXPOSE ${httpPort}

ENTRYPOINT ./server
