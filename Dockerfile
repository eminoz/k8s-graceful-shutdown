FROM golang:1.17-alpine3.15 as builder
WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .
EXPOSE 8080
RUN go build main.go


CMD ["./main"]