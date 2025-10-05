# backend
Бэкенд репозиторий сервиса такси

## Описание
Golang backend server для обработки API запросов для объектов Client, Driver и Car.

## Требования
- Go 1.16 или выше
- PostgreSQL 12 или выше

## Установка и запуск

### Установка зависимостей
```bash
go mod download
```

### Сборка
```bash
go build -o backend
```

### Запуск сервера
```bash
./backend
```

Сервер запустится на порту 8080 (по умолчанию).

### Переменные окружения

#### Основные настройки
- `SERVER_PORT` - порт сервера (по умолчанию: 8080)
- `DATABASE_URL` - полная строка подключения к PostgreSQL (опционально)

#### Настройки подключения к PostgreSQL (если DATABASE_URL не указан)
- `DB_HOST` - хост PostgreSQL (по умолчанию: localhost)
- `DB_PORT` - порт PostgreSQL (по умолчанию: 5432)
- `DB_USER` - имя пользователя PostgreSQL (по умолчанию: postgres)
- `DB_PASSWORD` - пароль пользователя PostgreSQL (по умолчанию: postgres)
- `DB_NAME` - имя базы данных (по умолчанию: taxi)
- `DB_SSLMODE` - режим SSL (по умолчанию: disable)

Примеры:
```bash
# Использование DATABASE_URL
DATABASE_URL="host=localhost port=5432 user=postgres password=secret dbname=taxi sslmode=disable" ./backend

# Использование отдельных переменных
DB_HOST=localhost DB_PORT=5432 DB_USER=postgres DB_PASSWORD=secret DB_NAME=taxi ./backend

# Минимальная конфигурация (используются значения по умолчанию)
./backend
```

### Настройка базы данных

Перед запуском сервера необходимо создать базу данных PostgreSQL:

```bash
# Подключиться к PostgreSQL
psql -U postgres

# Создать базу данных
CREATE DATABASE taxi;

# Выйти из psql
\q
```

Таблицы будут созданы автоматически при первом запуске сервера.

## API Endpoints

### Health Check
- `GET /health` - проверка состояния сервера

### Clients (Клиенты)

#### Получить всех клиентов
```bash
GET /api/clients
```

#### Получить клиента по ID
```bash
GET /api/clients/{id}
```

#### Создать нового клиента
```bash
POST /api/clients
Content-Type: application/json

{
  "name": "Ivan Petrov",
  "phone": "+79991234567",
  "email": "ivan@example.com"
}
```

#### Обновить клиента
```bash
PUT /api/clients/{id}
Content-Type: application/json

{
  "name": "Ivan Petrov Updated",
  "phone": "+79991234567",
  "email": "ivan.updated@example.com"
}
```

#### Удалить клиента
```bash
DELETE /api/clients/{id}
```

### Drivers (Водители)

#### Получить всех водителей
```bash
GET /api/drivers
```

#### Получить водителя по ID
```bash
GET /api/drivers/{id}
```

#### Создать нового водителя
```bash
POST /api/drivers
Content-Type: application/json

{
  "name": "Alexey Smirnov",
  "phone": "+79997654321",
  "license_number": "7712345678",
  "rating": 4.8
}
```

#### Обновить водителя
```bash
PUT /api/drivers/{id}
Content-Type: application/json

{
  "name": "Alexey Smirnov Updated",
  "phone": "+79997654321",
  "license_number": "7712345678",
  "rating": 4.9
}
```

#### Удалить водителя
```bash
DELETE /api/drivers/{id}
```

### Cars (Автомобили)

#### Получить все автомобили
```bash
GET /api/cars
```

#### Получить автомобиль по ID
```bash
GET /api/cars/{id}
```

#### Создать новый автомобиль
```bash
POST /api/cars
Content-Type: application/json

{
  "driver_id": 1,
  "brand": "Toyota",
  "model": "Camry",
  "year": 2020,
  "license_plate": "A123BC77",
  "color": "Black"
}
```

#### Обновить автомобиль
```bash
PUT /api/cars/{id}
Content-Type: application/json

{
  "driver_id": 1,
  "brand": "Toyota",
  "model": "Camry",
  "year": 2021,
  "license_plate": "A123BC77",
  "color": "White"
}
```

#### Удалить автомобиль
```bash
DELETE /api/cars/{id}
```

## Структура проекта
```
.
├── main.go              # Точка входа приложения
├── config/              # Конфигурация
│   └── config.go
├── models/              # Модели данных
│   ├── client.go
│   ├── driver.go
│   └── car.go
├── handlers/            # HTTP обработчики
│   ├── client.go
│   ├── driver.go
│   └── car.go
├── database/            # Работа с БД
│   └── database.go
├── go.mod
└── go.sum
```

## Примеры использования

### Создание клиента
```bash
curl -X POST http://localhost:8080/api/clients \
  -H "Content-Type: application/json" \
  -d '{"name":"Ivan Petrov","phone":"+79991234567","email":"ivan@example.com"}'
```

### Получение всех водителей
```bash
curl -X GET http://localhost:8080/api/drivers
```

### Обновление автомобиля
```bash
curl -X PUT http://localhost:8080/api/cars/1 \
  -H "Content-Type: application/json" \
  -d '{"driver_id":1,"brand":"Toyota","model":"Camry","year":2021,"license_plate":"A123BC77","color":"White"}'
```

## Документация

Этот проект включает полную документацию godoc для всех компонентов системы.

### Просмотр документации

#### 1. Документация через командную строку

Просмотр документации всех пакетов:
```bash
go doc -all
```

Просмотр документации конкретного пакета:
```bash
go doc config
go doc database
go doc models
go doc handlers
```

Просмотр документации конкретного типа:
```bash
go doc models.Client
go doc models.Driver
go doc models.Car
```

Просмотр документации конкретной функции:
```bash
go doc handlers.GetClients
go doc config.LoadConfig
go doc database.InitDB
```

#### 2. Веб-документация

Запуск локального сервера godoc:
```bash
godoc -http=:6060
```

Затем откройте браузер и перейдите по адресу: `http://localhost:6060/pkg/github.com/hse-trpo-taxi/backend/`

#### 3. Онлайн документация

Если проект доступен публично, документацию можно просмотреть онлайн по адресу:
`https://godocs.io/github.com/hse-trpo-taxi/backend`

### Структура документации

Проект включает полную документацию для:

- **Документация на уровне пакетов**: Обзор всей системы сервиса такси
- **Документация типов**: Подробные описания всех структур данных
- **Документация функций**: Полные описания всех публичных функций
- **Документация API**: Спецификации RESTful эндпоинтов
- **Документация конфигурации**: Описания переменных окружения
- **Документация схемы БД**: Структуры таблиц и связи

### Ключевые особенности документации

- **API эндпоинты**: Полная спецификация REST API с HTTP методами и кодами состояния
- **Руководство по конфигурации**: Переменные окружения и значения по умолчанию
- **Схема базы данных**: Структуры таблиц и связи
- **Примеры использования**: Примеры запросов и ответов
- **Обработка ошибок**: HTTP коды состояния и ответы об ошибках
- **Зависимости**: Необходимые внешние пакеты

### Стандарты документации

Вся документация соответствует соглашениям godoc Go:
- Комментарии пакетов начинаются с "Package [name]"
- Комментарии функций начинаются с имени функции
- Комментарии типов описывают назначение и использование
- Примеры предоставляются там, где это полезно
- HTTP коды состояния документированы для API эндпоинтов

### Генерация документации

Для повторной генерации документации после изменений кода:
```bash
go doc -all > documentation.txt
```

Документация автоматически генерируется из комментариев в исходном коде и остается актуальной вместе с кодовой базой.
