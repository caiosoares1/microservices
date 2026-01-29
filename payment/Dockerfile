FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY microservices-proto/ /microservices-proto/
COPY microservices/payment/ ./

RUN go mod edit \
    -replace github.com/ruandg/microservices-proto/golang/payment=/microservices-proto/golang/payment

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /payment ./cmd/main.go

FROM alpine:latest

WORKDIR /

COPY --from=builder /payment /payment

EXPOSE 50052

ENTRYPOINT ["/payment"]