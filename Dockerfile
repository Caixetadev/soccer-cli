FROM golang:1.20 as builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./

RUN go build -o /soccer-cli

ENTRYPOINT ["/soccer-cli"]