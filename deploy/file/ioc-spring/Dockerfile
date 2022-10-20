FROM golang:1.19.2

RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct

RUN go install github.com/alibaba/ioc-golang/iocli@latest
