FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.mod .
COPY main.go .
RUN go mod tidy
RUN go build -o weather_exporter

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/weather_exporter /app/weather_exporter
ENTRYPOINT /app/weather_exporter