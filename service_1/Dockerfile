FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o app main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/app .
EXPOSE 8000
CMD ["./app"]