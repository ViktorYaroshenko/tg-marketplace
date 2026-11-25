# tg-marketplace

Учебный микросервисный бэкенд маркетплейса на Go с интерфейсом в виде Telegram-бота.

## Сервисы

- `ledger`: счета, переводы с двойной записью, идемпотентность.
- `inventory`: остатки, резервирование и списание.
- `bot-gateway`: взаимодействие с пользователями через Telegram Bot API.

## Стек

Go, PostgreSQL, Redis, Kafka, Docker Compose, Prometheus, Grafana (позже Kubernetes).

## Запуск

Создать файл конфигурации:

```bash
cp deploy/.env.example deploy/.env
```

Заполнить необходимые переменные окружения в `deploy/.env`.

Запустить сервисы:

```bash
docker compose -f deploy/docker-compose.yml up -d --build
```

PostgreSQL доступен на `localhost:5433`, Kafka — на `localhost:9092`.

## Bot Gateway

Сервис взаимодействия с Telegram Bot API.

### Реализованный функционал

- Получение сообщений через long polling.
- Управление offset при получении обновлений.
- Обработка команды `/start`.
- Отправка сообщений пользователю.
- Обработка ошибок сетевых запросов.
- Graceful shutdown через context и сигналы SIGINT/SIGTERM.
- Многоэтапная сборка Docker-образа.
- Поддержка подключения через HTTP-прокси.

### Переменные окружения

| Переменная | Описание |
|---|---|
| `TELEGRAM_BOT_TOKEN` | Токен бота, полученный через BotFather |
| `BOT_HTTPS_PROXY` | Необязательный адрес прокси, доступный из контейнера |

При запуске через Docker Compose переменная `BOT_HTTPS_PROXY` передаётся в контейнер как `HTTPS_PROXY`.

При локальном запуске Go-приложения необходимо использовать переменную `HTTPS_PROXY` напрямую.

На текущем этапе бот поддерживает команду `/start` в личном чате. Интеграция с Kafka и другими сервисами будет реализована на следующих этапах.

### Запуск отдельно от других сервисов

```bash
docker compose -f deploy/docker-compose.yml up -d --build bot-gateway
```

### Просмотр логов

```bash
docker compose -f deploy/docker-compose.yml logs -f bot-gateway
```

### Остановка

```bash
docker compose -f deploy/docker-compose.yml stop bot-gateway
```

## Архитектура

Скоро: схема сервисов и событий.