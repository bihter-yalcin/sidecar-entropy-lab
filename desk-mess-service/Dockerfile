FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o desk-mess-service .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/desk-mess-service .

EXPOSE 9090

CMD ["./desk-mess-service"]