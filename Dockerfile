# Stage-1
FROM golang:1.17 AS builder

WORKDIR /build
COPY . .
RUN go get -d -v ./...

ENV CGO_ENABLED=0 GOOS=linux
RUN go build -o shorten-url .

# Stage-2
FROM alpine:3.13 AS certificates
RUN apk --no-cache add ca-certificates

# Stage-3
FROM scratch
#FROM golang:1.17.8-bullseye

WORKDIR /app
COPY --from=certificates /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /build/shorten-url .

EXPOSE 9090

CMD ["./shorten-url"]