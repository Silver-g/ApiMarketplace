FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o ./bin/app cmd/main.go

FROM alpine AS runner

COPY --from=builder /app/bin/app .
COPY docker.env .env

EXPOSE 8080

ENTRYPOINT ["/app"]