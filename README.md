# backend
Бэкенд репозиторий сервиса такси

## Описание
Golang backend server для обработки API запросов для объектов Client, Driver и Car.

## Требования
- Go 1.16 или выше
- SQLite3

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
- `SERVER_PORT` - порт сервера (по умолчанию: 8080)
- `DB_PATH` - путь к базе данных SQLite (по умолчанию: ./taxi.db)

Пример:
```bash
SERVER_PORT=3000 DB_PATH=/tmp/taxi.db ./backend
```

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
