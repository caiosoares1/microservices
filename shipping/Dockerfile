FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY microservices-proto/ /microservices-proto/
COPY microservices/shipping/ ./

RUN go mod edit \
    -replace github.com/caiosoares1/microservices-proto/golang/shipping=/microservices-proto/golang/shipping

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /shipping ./cmd/main.go

FROM alpine:latest

WORKDIR /

COPY --from=builder /shipping /shipping

EXPOSE 50053

ENTRYPOINT ["/shipping"]
