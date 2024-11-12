FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY main.go .

RUN go build -o server main.go

FROM gcr.io/distroless/static-debian12:latest

WORKDIR /

COPY --from=builder /app/server .

EXPOSE 8080

CMD ["./server"]
