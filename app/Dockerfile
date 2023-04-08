# マルチステージングビルド
# alpineは軽量版 こちらをデプロイする
FROM golang:1.19-alpine as prod
# FROM golang:1.19-alpine

RUN apk update && \
    apk add  curl && \
    apk add --no-cache git

# COPY ./app .
COPY ./ /app
# WORKDIR /go/src/app
WORKDIR /app

# Dockerfile上でgoモジュールインストールは無理そう
# RUN test -f /app/go.mod || go mod init app
# RUN go mod download
# RUN go mod tidy

EXPOSE 8080
CMD ["go", "run", "main.go"]

# ローカル開発環境用(ホットローディング)
FROM golang:1.19 as dev

RUN apt-get update && \
    apt-get install curl && \
    apt-get install -y git

RUN go install github.com/pilu/fresh@latest

# COPY ./app .
COPY ./ /app
# WORKDIR /go/src/app
WORKDIR /app

# 下記のコマンドはどちらでも動く
# CMD ["fresh"]
CMD ["fresh", "-c", "./fresh.conf"]

EXPOSE 8080

# ローカル開発環境用(ホットローディング + リモートデバッグ)
FROM golang:1.19 as dev-debug

RUN apt-get update && \
    apt-get install curl && \
    apt-get install -y git

# freshとgo-delveを使いリモートデバッグができなかった為、airを使用
RUN go install github.com/go-delve/delve/cmd/dlv@latest
RUN go install github.com/cosmtrek/air@latest

# COPY ./app .
COPY ./ /app
# WORKDIR /go/src/app
WORKDIR /app

CMD ["air", "-c", "./air.toml"]

EXPOSE 8080