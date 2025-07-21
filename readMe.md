# Marketplace API

REST API для маркетплейса на Go. Поддерживает регистрацию, авторизацию, размещение нового объявления, отображение ленты объявлений. База данных PostgreSQL.

## API Endpoints

### Регистрация
- **Эндпоинт**: `POST /register`
- **Тело запроса**:
  ```json
  {
      "login": "string",
      "password": "string"
  }
  ```

### Авторизация
- **Эндпоинт**: `POST /login`
- **Тело запроса**:
  ```json
  {
      "login": "string",
      "password": "string"
  }
  ```

### Создание объявления
- **Эндпоинт**: `POST /ads`
- **Тело запроса**:
  ```json
  {
      "title": "string",
      "description": "string",
      "image_url": "string",
      "price": decimal
  }
  ```

### Получение списка объявлений
- **Эндпоинт**: `GET /ads?page=(int)&limit=(int)&sort_by=price_low|price_high|created_at_new|created_at_old&price=<int(min)-int(max)`
- **Параметры запроса**:
  - `page`: Номер страницы.
  - `limit`: Количество записей на странице.
  - `sort_by`: Сортировка (`price_low`, `price_high`, `created_at_new`, `created_at_old`).
  - `price`: Диапазон цен (`min-max`, например, `10-100`).

## Установка и запуск

### Запуск в Docker
   ```bash
   docker compose --env-file docker.env up --build
   ```

### Локальный запуск
   ```bash
   go run ./cmd/main.go
   ```

## Миграции базы данных
- **Применение миграций**:
  ```bash
  migrate -path ./migrations -database "postgres://postgres:Admin@database:5432/marketplace_db?sslmode=disable" up
  
  docker-compose run --rm migrate sh -c "migrate -path /migrations -database 'postgres://postgres:Admin@database:5432/marketplace_db?sslmode=disable' -verbose up"
  ```
- **Откат миграций**:
  ```bash
  migrate -path ./migrations -database "postgres://postgres:Admin@database:5432/marketplace_db?sslmode=disable" down

  docker-compose run --rm migrate sh -c "migrate -path /migrations -database 'postgres://postgres:Admin@database:5432/marketplace_db?sslmode=disable' -verbose down -all"
  ```

## Примеры JSON
- **Регистрация/Авторизация**:
  ```json
  {
      "login": "TestInp",
      "password": "123456789"
  }
  ```
- **Создание объявления**:
  ```json
  {
      "title": "TestInp",
      "description": "TestInpTestInp",
      "image_url": "testUrl",
      "price": 12.12
  }
  ```

