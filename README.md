# 📚 Language Learning Platform

Backend-сервис для платформы изучения языков. Написан на Go, данные хранятся в PostgreSQL.

## Стек

- **Go** — стандартная библиотека `net/http`
- **PostgreSQL** — драйвер `lib/pq`

## Структура

```
├── main.go       # Точка входа и маршруты
├── language/     # Языки
├── users/        # Пользователи
├── progress/     # Прогресс обучения
└── pass.txt      # Пароль к БД (не коммитить!)
```

## Запуск

```bash
# Создать файл с паролем от PostgreSQL
echo "your_password" > pass.txt

# Установить зависимости и запустить
go mod download
go run main.go
```

Сервер поднимается на `http://localhost:8182`. Подключается к базе `LLP` под пользователем `postgres`.

## API

### /language
| Метод | URL | Действие |
|-------|-----|----------|
| GET | `/language` | Все языки |
| POST | `/language` | Добавить язык |
| GET | `/language/{id}` | Язык по ID |
| PUT/PATCH | `/language/{id}` | Обновить |
| DELETE | `/language/{id}` | Удалить |

### /users
| Метод | URL | Действие |
|-------|-----|----------|
| GET | `/users` | Все пользователи |
| POST | `/users` | Создать пользователя |
| GET | `/users/{id}` | Пользователь по ID |
| PUT/PATCH | `/users/{id}` | Обновить |
| DELETE | `/users/{id}` | Удалить |

### /progress
| Метод | URL | Действие |
|-------|-----|----------|
| GET | `/progress` | Весь прогресс |
| GET | `/progress/{id}` | Прогресс по ID |
| PUT/PATCH | `/progress/{id}` | Обновить прогресс |
