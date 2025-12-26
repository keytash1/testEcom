# Todo Manager

HTTP-сервер для управления задачами на Golang. Тестовое задание для ecom.tech

# API

| POST | `/todos` | Создать новую задачу | 201, 400 |

| GET | `/todos` | Получить все задачи | 200 |

| GET | `/todos/{id}` | Получить задачу по id | 200, 400, 404 |

| PUT | `/todos/{id}` | Обновить задачу (полная замена) | 200, 400, 404 |

| DELETE | `/todos/{id}` | Удалить задачу по id | 204, 400, 404 |

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
curl -X POST http://localhost:8080/todos \
  -H "Content-Type: application/json" \
  -d '{"title":"Пройти стажировку в ecom.tech","description":"Важно!"}'

# Получить все
curl http://localhost:8080/todos

# Получить по ID  
curl http://localhost:8080/todos/1

# Обновить задачу (заменить)
curl -X PUT http://localhost:8080/todos/1 \
  -H "Content-Type: application/json" \
  -d '{"title":"Перейти в штат ecom.tech","description":"Очень важно!","completed":true}'

# Удалить
curl -X DELETE http://localhost:8080/todos/1

# Тестирование
go test ./...