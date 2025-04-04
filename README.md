# Warehouse Management Service

# Микросервис для управления складами, товарами и аналитикой продаж


## 📌 Техническое задание

Реализация HTTP-сервиса для:
- Управления складами (warehouses)
- Управления товарами (products)
- Инвентаризации (inventory)
- Аналитики продаж (analytics)
- Обработки покупок


## 🛠 Технологический стек

- **Язык**: Go 1.24
- **База данных**: PostgreSQL 16+
- **Драйвер БД**: pgx/v5
- **Миграции**: go-migrate
- **Логирование**: zap logger
- **Контейнеризация**: Docker
- **Документация**: Swagger

## 📦 Установка и запуск

### 1. Локальный запуск (без Docker)

```bash
# Клонирование репозитория
git clone https://github.com/SamuraiAkira/warehouse-management-service.git
cd warehouse-management-service

# Установка зависимостей
go mod download

# Настройка БД (требуется пароль от PostgreSQL)
createdb warehouse -U postgres

# Применение миграций
migrate -path migrations -database "postgres://postgres:your_password@localhost:5432/warehouse?sslmode=disable" up

# Запуск сервера
go run cmd/app/main.go
```

### 2. Запуск с Docker

```bash
docker-compose up --build
```

Сервис будет доступен на `http://localhost:8080`

## 🌐 API Endpoints

Доступные API endpoints:

- `GET    /api/health` - Health check сервиса
- `GET    /api/warehouses` - Список складов
- `POST   /api/warehouses` - Создание склада
- `GET    /api/products` - Список товаров
- `POST   /api/products` - Создание товара
- `PUT    /api/products/{id}` - Обновление товара
- `POST   /api/inventory` - Добавление товара на склад
- `POST   /api/sales` - Оформление продажи

Полная документация: [Swagger UI](http://localhost:8080/swagger/index.html)

## 🏗 Структура проекта

```
.
├── cmd/app/               # Главный пакет приложения
├── internal/              # Внутренние пакеты
│   ├── app/               # Ядро приложения
│   │   ├── config/        # Конфигурация
│   │   ├── entity/        # Сущности
│   │   ├── repository/    # Репозитории
│   │   ├── service/       # Бизнес-логика
│   │   └── delivery/      # Доставка (HTTP handlers)
├── migrations/            # SQL миграции
├── pkg/                   # Вспомогательные пакеты
├── deployments/           # Docker файлы
├── docs/                  # Документация
└── scripts/               # Вспомогательные скрипты
```

## 🧪 Тестирование

```bash
# Unit-тесты
go test ./...

# Интеграционные тесты (требует запущенную БД)
go test -tags=integration ./...
```

## 🚀 Развертывание

Production-сборка:
```bash
docker build -t warehouse-service -f deployments/Dockerfile.prod .
```

