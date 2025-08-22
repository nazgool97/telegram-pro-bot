# Этап сборки
FROM golang:1.23-alpine AS builder
WORKDIR /app

# Модули
COPY go.mod go.sum ./
RUN go mod download

# Копируем весь исходник
COPY . .

# Собираем бинарник из cmd/main.go
RUN CGO_ENABLED=0 go build -o server ./cmd/main.go

# Финальный образ
FROM alpine:3.19
WORKDIR /root/
COPY --from=builder /app/server .
EXPOSE 8080
CMD ["./server"]
