# gRPC Demo

Демонстрационный проект gRPC сервера и клиента на Go с реализацией graceful shutdown.

## Описание

Этот проект демонстрирует создание простого gRPC сервиса с использованием Go. Проект включает:

- **gRPC сервер** с Echo сервисом
- **gRPC клиент** для тестирования
- **Graceful shutdown** с таймаутом
- **Конфигурация через переменные окружения**
- **Protobuf определения**

## Структура проекта

```
grpc-demo/
├── cmd/
│   ├── client/          # gRPC клиент
│   └── server/          # gRPC сервер
├── echo/                # Сгенерированные protobuf файлы
├── pkg/
│   └── helper.go        # Вспомогательные функции
├── proto/
│   └── echo.proto       # Protobuf определения
├── go.mod
├── go.sum
└── README.md
```

## Требования

- Go 1.23.4 или выше
- Protocol Buffers compiler (protoc)

## Установка и настройка

### 1. Клонирование репозитория

```bash
git clone https://github.com/akozadaev/grpc-demo.git
cd grpc-demo
```

### 2. Установка зависимостей

```bash
go mod download
```

### 3. Генерация protobuf файлов

```bash
 protoc --go_out=. --go-grpc_out=. ./proto/echo.proto

```

### 4. Создание файла конфигурации

Создайте файл `.env` в корне проекта:

```env
NETWORK=tcp
PORT=50051
```

## Запуск

### Запуск сервера

```bash
go run cmd/server/main.go
```

Сервер запустится на порту, указанном в переменной окружения `PORT` (по умолчанию 50051).

### Запуск клиента

```bash
go run cmd/client/main.go
```

## API

### EchoService

Сервис предоставляет простой метод Echo для эхо-ответов.

#### Echo

**Запрос:**
```protobuf
message EchoRequest {
  string message = 1;
}
```

**Ответ:**
```protobuf
message EchoResponse {
  string message = 1;
}
```

**Описание:** Возвращает входящее сообщение с префиксом "Echo: ".

## Graceful Shutdown

Сервер реализует корректное завершение работы:

1. **Ожидание сигналов:** `SIGINT` (Ctrl+C) или `SIGTERM`
2. **Graceful shutdown:** Прекращение приема новых соединений
3. **Таймаут:** 5 секунд на завершение текущих запросов
4. **Fallback:** Принудительное завершение при превышении таймаута

### Пример завершения:

```bash
# Запуск сервера
go run cmd/server/main.go

# В другом терминале - отправка сигнала
kill -TERM <PID>
```

Вывод при graceful shutdown:
```
gRPC сервер слушает на :50051
Начинается graceful shutdown...
gRPC сервер успешно завершен
```

## Переменные окружения

| Переменная | Описание | Значение по умолчанию |
|------------|----------|----------------------|
| `NETWORK`  | Тип сети для gRPC сервера | `tcp` |
| `PORT`     | Порт для gRPC сервера | `50051` |

## Зависимости

- `google.golang.org/grpc` - gRPC framework
- `google.golang.org/protobuf` - Protocol Buffers
- `github.com/joho/godotenv` - Загрузка переменных окружения

## Разработка

### Добавление новых методов

1. Обновите `proto/echo.proto`
2. Сгенерируйте новые файлы: `protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/echo.proto`
3. Реализуйте методы в `cmd/server/main.go`

### Тестирование

Для тестирования можно использовать:
- Встроенный клиент: `go run cmd/client/main.go`
- gRPC CLI инструменты (grpcurl, grpc_cli)
- Postman с поддержкой gRPC

## Лицензия

MIT License
