# 📎 ShortURLS

Простой сервис сокращения ссылок на Go с использованием Echo и PostgreSQL.

## 🚀 Возможности

- Сокращение длинных ссылок в уникальные ключи
- Редирект по короткому URL
- REST API
- Хранение в PostgreSQL

## 🧱 Стек

- Go
- Echo
- PostgreSQL
- godotenv

## 📦 Установка

1. Клонируйте репозиторий:

```bash
git clone https://github.com/Wucop228/ShortURLS.git
cd ShortURLS
```

2. Создайте env и заполните его для данных о бд
```bash
touch .env
```
```
DB_USER=your_user
DB_PASSWORD=your_password
DB_NAME=your_db
DB_PORT=5432
DB_HOST=localhost
PORT=8080
```

3. Запустите сервер
```bash
go run cmd/server/main.go
```

4. Примеры использования API

**Пример создания через curl-запрос**
```
curl -X POST http://localhost:8080/create-url \
  -H "Content-Type: application/json" \
  -d '{"url": "https://example.com"}'
```

**Пример использования**
```
curl -v http://localhost:8080/urls/example
```
или можно увидеть перенаправление в браузере, введя эту ссылку http://localhost:8080/urls/example