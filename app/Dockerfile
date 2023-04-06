# マルチステージングビルド

# ローカル開発環境用(ホットローディング)
FROM golang:1.19 as dev

WORKDIR /go/src/app
# COPY ./app .
COPY ./ /app

RUN apt-get update && \
    apt-get install curl && \
    apt-get install -y git

RUN go install github.com/pilu/fresh@latest

# 下記のコマンドはどちらでも動く
# CMD ["fresh"]
CMD ["fresh", "-c", "./fresh.conf"]

EXPOSE 8080

# ローカル開発環境用(ホットローディング + リモートデバッグ)
FROM golang:1.19 as dev-debug

WORKDIR /go/src/app
# COPY ./app .
COPY ./ /app

RUN apt-get update && \
    apt-get install curl && \
    apt-get install -y git

# freshとgo-delveを使いリモートデバッグができなかった為、airを使用
RUN go install github.com/go-delve/delve/cmd/dlv@latest
RUN go install github.com/cosmtrek/air@latest

CMD ["air", "-c", "./air.toml"]

EXPOSE 8080

# alpineは軽量版 こちらをデプロイする
# FROM golang:1.19-alpine as prod
FROM golang:1.19-alpine

WORKDIR /go/src/app
# COPY ./app .
COPY ./ /app

RUN apk update && \
    apk add  curl && \
    apk add --no-cache git

EXPOSE 8080

CMD ["go", "run", "main.go"]