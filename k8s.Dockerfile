# 第一層基底
FROM golang:1.11.2-alpine AS build

# 載入翻譯包
RUN apk add git \
    && go get github.com/liuzl/gocc

# 複製原始碼
COPY . /go/src/dragon-fruit
WORKDIR /go/src/dragon-fruit

# 進行編譯(名稱為：dragon-fruit)
RUN go build -o dragon-fruit

# 最終運行golang 的基底
FROM alpine

#  AWS 需要此套件
RUN apk update \
    && apk add ca-certificates \
    && rm -rf /var/cache/apk/*

COPY --from=build /go/src/dragon-fruit/dragon-fruit /app/dragon-fruit

## 複製所需要的檔案
COPY ./env /app/env
COPY ./goeip /app/goeip


WORKDIR /app

# 設定容器時區(美東)
RUN apk update \
    && apk add tzdata \
    && cp /usr/share/zoneinfo/America/Puerto_Rico /etc/localtime

RUN mkdir -p /app/log/
RUN ln -sf /dev/stdout /app/log/dragon-fruit_access.log \
    && ln -sf /dev/stdout /app/log/dragon-fruit_error.log

ENTRYPOINT [ "./dragon-fruit" ]