FROM golang:1.19.2-alpine3.16

RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct

WORKDIR /go/server
COPY . .

RUN echo -e "https://mirror.tuna.tsinghua.edu.cn/alpine/v3.4/main\n\ https://mirror.tuna.tsinghua.edu.cn/alpine/v3.4/community" > /etc/apk/repositories
RUN apk update
RUN apk add gcc g++ make cmake gfortran libffi-dev openssl-dev libtool

RUN go install github.com/swaggo/swag/cmd/swag

RUN swag init --pd

RUN go build -v -mod=vendor -o server .

FROM alpine:3.16.2

WORKDIR /go/server

COPY --from=0 /go/server/server ./
COPY --from=0 /go/server/configs ./configs/

ENTRYPOINT ./server
