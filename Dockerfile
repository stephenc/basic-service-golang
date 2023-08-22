FROM golang:1.21-alpine AS builder

WORKDIR /app/server

ENV CGO_ENABLED=0

COPY go* .

RUN go mod download

COPY . .

RUN go install -v

FROM alpine:latest

EXPOSE 8080

COPY --from=builder /go/bin/basic-service-golang /

ENTRYPOINT /basic-service-golang
