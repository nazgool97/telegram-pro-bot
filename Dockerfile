# syntax=docker/dockerfile:1

FROM golang:1.23 AS builder

WORKDIR /app

# Копируем модули и скачиваем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем все файлы проекта
COPY . .

# Собираем бинарник
RUN go build -o server ./cmd/main.go

# Финальный образ
FROM gcr.io/distroless/base-debian12:latest
WORKDIR /root/
COPY --from=builder /app/server .

CMD ["./server"]
