# Etapa de construcción
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Etapa de ejecución
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
COPY db/ /app/db/
EXPOSE 8080
CMD ["./main"]