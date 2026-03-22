FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o news-zve .

FROM alpine:3.19
RUN apk add --no-cache ca-certificates

WORKDIR /root/

COPY --from=builder /app/news-zve .
COPY --from=builder /app/html ./html
COPY --from=builder /app/media ./media

EXPOSE 8080

CMD ["./news-zve"]