FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY microservices-proto/ /microservices-proto/
COPY microservices/order/ ./

RUN go mod edit \
    -replace github.com/caiosoares1/microservices-proto/golang/order=/microservices-proto/golang/order \
    -replace github.com/caiosoares1/microservices-proto/golang/payment=/microservices-proto/golang/payment \
    -replace github.com/caiosoares1/microservices-proto/golang/shipping=/microservices-proto/golang/shipping

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /order ./cmd/main.go

FROM alpine:latest

WORKDIR /

COPY --from=builder /order /order

EXPOSE 50051

ENTRYPOINT ["/order"]
