FROM golang:1.12 as builder
LABEL maintainer="Fajar AR <github.com/farkroft>"
WORKDIR /app
RUN go mod init /app
RUN go mod vendor
COPY . .
COPY application.yaml /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod vendor -ldflags="-w -s" -o auth-service

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root
COPY --from=builder /app/auth-service .
EXPOSE 8080
CMD ["./auth-service"]