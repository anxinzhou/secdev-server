FROM golang:1.10

WORKDIR /app

COPY . .

ENV GOPATH=/app

WORKDIR /app/src

CMD ["go","run","server.go"]