# Warehouse Management Service

Микросервис для управления складами, товарами.

![Go Version](https://img.shields.io/badge/Go-1.24+-blue)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-16+-blue)
![License](https://img.shields.io/badge/License-MIT-green)

## 📌 Особенности

- Управление складами 
- Управление товарами с характеристиками
- Инвентаризация с учетом остатков и цен
- Аналитика продаж по складам
- REST API
- Чистая архитектура и SOLID принципы
- Graceful shutdown сервера
- Логирование запросов с X-Request-ID

## ВАЖНО

Если вы видите этот текст значит я не закончил работу и часть функцианала не работает!

### Требования
- Docker 20.10+
- Docker Compose 2.0+
- Go 1.24 (для локальной разработки)

### Запуск через Docker
```
git clone https://github.com/SamuraiAkira/warehouse-management-service.git
cd warehouse-management-service
docker-compose up -d
```
Сервис будет доступен на http://localhost:8080

📚 API Endpoints
Склады
```
POST /api/warehouses - Создать склад
```
```
GET /api/warehouses - Список складов
```
Товары
```
POST /api/products - Создать товар
```
```
GET /api/products - Список товаров
```
Инвентаризация:
```
POST /api/inventory - Добавить товар на склад
```
```
GET /api/inventory/{warehouse_id} - Товары на складе
```

🛠 Разработка
Структура проекта
```
.
├── cmd/app/          # Точка входа
├── internal/         # Внутренние пакеты
│   ├── app/          # Ядро приложения
│   ├── pkg/          # Вспомогательные пакеты
├── migrations/       # SQL-миграции
├── deployments/      # Конфигурации развертывания
```
Локальный запуск
```
go run cmd/app/main.go
```
Тестирование
```
go test ./...
```
🔧 Настройка
Конфигурация через .env файл:
```
DB_HOST=postgres
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=warehouse
```