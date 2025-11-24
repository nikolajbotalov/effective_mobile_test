# Subscriptions Service

Subscriptions Service - это сервис для агрегации данных об онлайн-подписках пользователей. Приложение позволяет создавать, обновлять, удалять и получать информацию о подписках, а также рассчитывать общую стоимость подписок пользователя за определенный период.

## Архитектура

Проект построен по принципам Clean Architecture с использованием следующих слоев:

- **cmd** - точка входа в приложение
- **internal/app** - основная логика приложения
- **internal/domain** - бизнес-сущности и модели
- **internal/handlers** - обработчики HTTP-запросов
- **internal/usecases** - бизнес-логика (use cases)
- **internal/repositories** - взаимодействие с базой данных
- **internal/adapters/db** - адаптер базы данных
- **docs** - документация API (Swagger)

## Технологии

- **Go** (версия 1.25.1) - основной язык программирования
- **PostgreSQL** - реляционная база данных
- **Gin** - веб-фреймворк
- **Swaggo** - генерация документации API
- **Docker & Docker Compose** - контейнеризация
- **golang-migrate** - миграции базы данных
- **Zap** - логирование

## Установка и запуск

### Локальный запуск

1. Установите Go (версия 1.25.1 или выше)
2. Установите PostgreSQL
3. Клонируйте репозиторий:
   ```bash
   git clone https://github.com/nikolajbotalov/effective_mobile_test
   cd effective_mobile_test
   ```
4. Создайте файл `.env` на основе `.env.example` и укажите настройки базы данных
5. Установите зависимости:
   ```bash
   go mod download
   ```
6. Запустите приложение:
   ```bash
   go run cmd/app/main.go
   ```

### Запуск через Docker

1. Убедитесь, что у вас установлен Docker и Docker Compose
2. Клонируйте репозиторий:
   ```bash
   git clone https://github.com/nikolajbotalov/effective_mobile_test
   cd effective_mobile_test
   ```
3. Создайте файл `.env` на основе `.env.example`
4. Запустите сервисы:
   ```bash
   docker-compose up --build
   ```

Приложение будет доступно по адресу `http://localhost:8080`

## API Endpoints

API документирован с использованием Swagger. Документация доступна по пути `/swagger/index.html` при запущенном приложении.

### Основные endpoints:

- `GET /api/v1/subscriptions` - получить список подписок с пагинацией
- `POST /api/v1/subscriptions` - создать новую подписку
- `GET /api/v1/subscriptions/{id}` - получить подписку по ID
- `PUT /api/v1/subscriptions/{id}` - обновить подписку
- `DELETE /api/v1/subscriptions/{id}` - удалить подписку
- `GET /api/v1/subscriptions/total_cost/{user_id}` - получить общую стоимость подписок пользователя за период

### Формат даты

Для полей даты используется формат `YYYY-MM-DD` (например, `2025-01-15`).

### Параметры фильтрации

Для получения списка подписок поддерживаются параметры пагинации:

- `limit` - количество записей на странице (по умолчанию 10)
- `offset` - смещение (по умолчанию 0)

Для получения общей стоимости поддерживаются параметры:

- `start_date` - начало периода (в формате MM-YYYY, например, 01-2025)
- `end_date` - конец периода (в формате MM-YYYY)
- `service_name` - фильтр по названию сервиса (опционально)

## Структура подписки

```go
type Subscription struct {
    ID          string     `json:"id"`
    ServiceName string     `json:"service_name"`
    Price       uint       `json:"price"`
    UserID      string     `json:"user_id"`
    StartDate   time.Time `json:"start_date"`
    EndDate     *time.Time `json:"end_date"`
    CreatedAt   time.Time  `json:"created_at"`
    UpdatedAt   time.Time  `json:"updated_at"`
}
```

## База данных

Проект использует PostgreSQL в качестве основной базы данных. Миграции находятся в папке `migrations/` и автоматически применяются при запуске приложения.

## Логирование

Приложение использует библиотеку Zap для логирования. Логи выводятся в формате JSON.

## Валидация

Для валидации входных данных используется встроенная валидация Gin и кастомные валидаторы. Проверяются:

- Обязательные поля
- Формат дат
- Корректность UUID
- Ограничения на значения (например, цена должна быть больше 0)

## Миграции

Миграции базы данных управляются с помощью библиотеки `golang-migrate`. При запуске приложения миграции применяются автоматически.

## Переменные окружения

Для корректной работы приложения необходимо настроить следующие переменные окружения:

```
BIND_IP=0.0.0
PORT=8080
PG_USER=postgres
PG_PASSWORD=admin
PG_HOST=db
PG_PORT=5432
PG_NAME=subscriptions
RETRY_ATTEMPTS=10
DATABASE_URL=postgres://postgres:admin@db:5432/subscriptions?sslmode=disable
POSTGRES_USER=postgres
POSTGRES_PASSWORD=admin
POSTGRES_DB=subscriptions
```
