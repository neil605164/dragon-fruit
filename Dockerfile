# 第一層基底
FROM golang:1.11.2-alpine

# 安裝 git
# go get fresh, grpc
RUN apk add git \
    && go get github.com/pilu/fresh \
    && go get google.golang.org/grpc