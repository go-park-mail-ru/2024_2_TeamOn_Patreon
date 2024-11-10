# Сборка docker контейнера
# >> docker build -t auth -f cmd/microservices/auth/Dockerfile .
# Запуск докер контейнера
# >> docker run -p 8080:8080  -v /absolute/path/cmd/microservices/auth/volume:/app/volume auth

# Еще залогиниться надо
# >> docker login

# Используем официальный образ Go для сборки приложения
FROM golang:1.23-alpine AS builder

# Устанавливаем зависимости для работы с файлом env.docker
RUN apk update && apk add --no-cache git

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем go.mod и go.sum для загрузки зависимостей
COPY go.mod go.sum ./

# Загружаем зависимости
RUN go mod tidy

# Копируем все файлы проекта, начиная с корня
COPY . .

# Собираем бинарный файл Go
RUN go build -o /app/auth-service cmd/microservices/auth/main.go

# Создаём финальный образ
FROM alpine:3.18

# Устанавливаем зависимости для работы с файлами env.docker
RUN apk update && apk add --no-cache bash

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем скомпилированное приложение из builder-контейнера
COPY --from=builder /app/auth-service /app/auth-service

# Копируем файл с переменными окружения из директории pkg
#COPY internal/auth/pkg/env.docker /app/env

# Экспортируем переменные окружения из файла env.docker
# Используем формат key=value
# RUN export $(grep -v '^#' /app/env | xargs)

# Пробрасываем порт 8080 снаружи на 8081 внутри контейнера
EXPOSE 8081

# Запускаем приложение
CMD ["sh", "-c", "/app/auth-service --port=8081"]
