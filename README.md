# SimpleCalc

**SimpleCalc** – это веб-сервис для вычисления арифметических выражений. Поддерживаются базовые арифметические операции, работа со скобками, отрицательные числа и числа с плавающей точкой.

---

## Функциональность

Сервис принимает арифметическое выражение через HTTP POST-запрос и возвращает результат вычислений.


---
### Поддерживаемые функции:

#### Успешные вычисления

| **Описание**               | **Выражение**        | **Ожидаемый результат** |
|----------------------------|----------------------|-------------------------|
| Простое сложение           | `1 + 2`              | `3`                     |
| Простое умножение          | `2 * 3`              | `6`                     |
| Простое деление            | `4 / 2`              | `2`                     |
| Простое вычитание          | `10 - 3`             | `7`                     |
| Смешанные операторы        | `3 + 5 * 2`          | `13`                    |
| Скобочный приоритет        | `(3 + 5) * 2`        | `16`                    |
| Отрицательные числа        | `-5 + 3`             | `-2`                    |
| Отрицательные в скобках    | `(-5 + 3) * 2`       | `-4`                    |
| Вложенные скобки           | `(1 + (2 * 3))`      | `7`                     |
| Двойные скобки             | `((2 + 3) * 4)`      | `20`                    |
| Числа с плавающей точкой   | `2.5 + 3.1`          | `5.6`                   |
| Деление с плавающей точкой | `5 / 2`              | `2.5`                   |
| Пробелы в выражении        | ` 2 + 3 `            | `5`                     |

#### Ошибки

| **Описание**               | **Выражение**        | **Ошибка** | **Тип ошибки**                |
|----------------------------|----------------------|------------|-------------------------------|
| Неверные символы           | `2 + *`              | `true`     | `errors.ErrInvalidExpression` |
| Несогласованные скобки     | `(1 + 2`             | `true`     | `errors.ErrInvalidExpression` |
| Пустое выражение           |                      | `true`     | `errors.ErrInvalidExpression` |
| Недопустимый символ        | `2 + @3`             | `true`     | `errors.ErrInvalidCharacter`  |
| Деление на ноль            | `10 / 0`             | `true`     | `errors.ErrDivideByZero`      |





---

## Примеры использования

### Успешные запросы (200 OK)

#### Простое сложение
**Запрос:**
```bash
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{"expression": "1+2"}'
```

**Ответ:**
```json
{
    "result": "3"
}
```

#### Умножение
**Запрос:**
```bash
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{"expression": "2*3"}'
```

**Ответ:**
```json
{
    "result": "6"
}
```

#### Скобки
**Запрос:**
```bash
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{"expression": "(3+5)*2"}'
```

**Ответ:**
```json
{
    "result": "16"
}
```

---

### Ошибки клиента (422 Unprocessable Entity)

#### Некорректное выражение
**Запрос:**
```bash
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{"expression": "2+*"}'
```

**Ответ:**
```json
{
    "error": "Expression is not valid"
}
```

#### Деление на ноль
**Запрос:**
```bash
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{"expression": "10/0"}'
```

**Ответ:**
```json
{
    "error": "Expression is not valid"
}
```

---

## Установка и запуск

### Требования:
- **Go** версии 1.19 или выше.

### Установка:
1. Клонируйте репозиторий:
   ```bash
   git clone https://github.com/w0ikid/simplecalc.git
   cd simplecalc
   ```

2. Установите зависимости:
   ```bash
   go mod tidy
   ```

3. Запустите сервер:
   ```bash
   go run ./cmd/calculator/main.go
   ```

Сервер будет доступен по адресу: `http://localhost:8080`.

---

## Тестирование

### Запуск всех тестов:
```bash
go test ./...
```

## Структура проекта

```plaintext
.
├── cmd
│   └── calculator        # Основной файл для запуска сервера
│       └── main.go
├── internal
│   ├── api               # HTTP-обработчики
│   │   ├── handler.go
│   │   └── handler_test.go
│   └── service           # Логика вычислений
│       ├── service.go
│       └── service_test.go
├── pkg
│   └── errors            # Определение ошибок
│       └── errors.go
├── go.mod                # Модуль Go
├── go.sum                # Контроль зависимостей
└── README.md             # Документация
```

---

## Технологии

- **Язык программирования:** Go.
- **Архитектура:** модульная, с разделением HTTP-обработки и бизнес-логики.

---

## Контакты

- **GitHub:** [https://github.com/w0ikid](https://github.com/w0ikid)