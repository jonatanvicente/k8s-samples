FROM golang:1.13 AS builder

WORKDIR /app
COPY main.go .
RUN CGO_ENABLED=0 GOOS=linux GOPROXY=https://proxy.golang.org go build -o app ./main.go

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/app .
CMD ["./app"]