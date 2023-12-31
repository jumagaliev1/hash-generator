FROM golang:latest AS builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN GO111MODULE="on" CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app ./cmd/api

FROM alpine:latest
WORKDIR /app
COPY --from=0 /app/app .
COPY --from=0 /app/internal/config /app/internal/config
CMD ["./app"]