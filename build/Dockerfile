FROM golang:1.23.3 AS builder
COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o /bin/app ./cmd/

FROM alpine:latest
COPY --from=builder /app/config /config
COPY --from=builder /bin/app /app

EXPOSE 8080
CMD ["/app"]
