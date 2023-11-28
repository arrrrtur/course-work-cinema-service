# Установка образа для сборки
FROM golang:1.17-alpine AS build

# Установка рабочей директории
WORKDIR /newApp

# Копирование файлов проекта в контейнер
COPY ./newApp ./

# Сборка приложения
RUN go build -o main .

# Установка образа для запуска
FROM alpine:latest

# Установка рабочей директории
WORKDIR /newApp

# Копирование бинарного файла из образа сборки в текущий образ
COPY --from=build /newApp/main .

# Определение команды запуска приложения
CMD ["./main"]
