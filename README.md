# Go URL Downloader CLI

CLI утилита на Golang для параллельной загрузки данных по списку URL с ограничением количества одновременных запросов (worker pool).

---

## Возможности

- Загрузка списка URL:
  - из аргументов командной строки
  - из файла
- Параллельная обработка (goroutines)
- Ограничение количества воркеров (worker pool)
- Таймауты для HTTP-запросов
- Вывод результата по каждому URL:
  - статус код
  - размер ответа
  - время выполнения
- Обработка ошибок (без падения приложения)

---

## Технологии

- Go (Golang)
- Стандартная библиотека (net/http, context, sync)

---

## Установка

```bash
git clone https://github.com/yourusername/go-url-downloader.git
cd go-url-downloader
go mod tidy
```

---

## Использование

1. Передача URL через аргументы

```bash
go run cmd/cli/main.go https://example.com https://google.com
```
---
2. Использование файла

Создай файл urls.txt:

```
https://example.com
https://google.com
```

Запуск:

```bash
go run cmd/cli/main.go -f urls.txt
```

---

## Параметры

| Флаг | Описание              | По умолчанию |
| ---- | --------------------- | ------------ |
| `-w` | Количество воркеров   | 5            |
| `-f` | Путь к файлу с URL    | —            |
| `-t` | Таймаут запроса (сек) | 5            |

---

Пример:

```bash
go run cmd/cli/main.go -w 10 -t 3 -f urls.txt
```

---

## Пример вывода

```
[OK]  https://example.com 200 15KB 120ms
[ERR] https://badsite.com timeout
```

---

## Как это работает

Приложение использует паттерн Worker Pool:

1. Список URL преобразуется в задачи (jobs)
2. Создаётся пул воркеров (goroutines)
3. Воркеры читают задачи из канала
4. Выполняют HTTP-запрос
5. Отправляют результат в канал результатов
6. Главная горутина собирает и выводит результаты

---

## Структура проекта

```
.
├── cmd/
│   └── cli/
│       └── main.go
├── internal/
│   ├── downloader/
│   │   ├── worker.go
│   │   └── client.go
│   ├── model/
│   │   └── models.go
│   └── config/
│       └── config.go
├── go.mod
└── README.md
```

---

## Основные структуры

```go
type Job struct {
    URL string
}

type Result struct {
    URL        string
    StatusCode int
    Size       int
    Duration   time.Duration
    Err        error
}
```

---

## Важные детали реализации

- Используется один http.Client (повышает производительность)
- Все ответы закрываются (defer resp.Body.Close())
- Контроль завершения через sync.WaitGroup
- Используется context.WithTimeout для запросов

---

## Возможные улучшения

- Retry механизм
- Rate limiting (ограничение RPS)
- Сохранение результата в файл (-o output.json)
- Прогресс-бар
- Unit-тесты

---

## Пример команды для разработки

```bash
go run cmd/cli/main.go -w 5 -t 5 https://example.com https://google.com
```

---

## Цель проекта
Этот проект создан как учебный, но с упором на production-подход:

- понимание concurrency в Go
- работа с HTTP
- реализация worker pool
- написание чистого и поддерживаемого кода

---

## Лицензия

MIT

---