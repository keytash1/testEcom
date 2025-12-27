# Todo Manager

HTTP-сервер для управления задачами на Golang. Тестовое задание для ecom.tech

# API

| POST | `/todos` | Создать новую задачу | 201, 400 |

| GET | `/todos` | Получить все задачи | 200 |

| GET | `/todos/{id}` | Получить задачу по id | 200, 400, 404 |

| PUT | `/todos/{id}` | Обновить задачу (полная замена) | 200, 400, 404 |

| DELETE | `/todos/{id}` | Удалить задачу по id | 204, 400, 404 |

| PATCH | `/todos/{id}/complete` | Изменить статус завершенности задачи | 200, 400, 404 |

# Доп. фича

PATCH ручка чтобы можно было удобно изменить статус задачи по id, просто передав true или false.

# Сервер запустится на http://localhost:8080
Клонировать репозиторий:

git clone https://github.com/keytash1/testEcom

cd todos_manager

# Запуск в докере
docker build -t todo-manager .

docker run -p 8080:8080 todo-manager

# Запуск без докера
go run ".\cmd\server\main.go"


# Пример использования API
# Создать
curl -X POST http://localhost:8080/todos -H "Content-Type: application/json" -d '{"title":"Complete internship at ecom.tech","description":"Important!"}'

# Получить все
curl http://localhost:8080/todos

# Получить по ID
curl http://localhost:8080/todos/1

# Обновить задачу (заменить)
curl -X PUT http://localhost:8080/todos/1 -H "Content-Type: application/json" -d '{"title":"Get hired at ecom.tech","description":"Very important!","completed":true}'

# Отметить задачу как выполненную
curl -X PATCH http://localhost:8080/todos/1/complete -H "Content-Type: application/json" -d '{"completed": true}'

# Снять отметку выполненной
curl -X PATCH http://localhost:8080/todos/1/complete -H "Content-Type: application/json" -d '{"completed": false}'

# Удалить
curl -X DELETE http://localhost:8080/todos/1

# Тестирование
go test ./...

P.S. В ТЗ сказано про тесты на дубликат id, но в моей реализации это невозможно т.к. не юзер присылает id, а "база данных" генерирует