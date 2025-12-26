"" 
docker build -t todo-manager .
docker run -p 8080:8080 todo-manager
go run ".\cmd\server\main.go"


# Создать
curl.exe -X POST http://localhost:8080/todos -H "Content-Type: application/json" -d '{\"title\":\"test\",\"description\":\"test\"}'

# Получить все
curl.exe http://localhost:8080/todos

# Получить по ID  
curl.exe http://localhost:8080/todos/1

# Обновить
curl.exe -X PUT http://localhost:8080/todos/1 -H "Content-Type: application/json" -d '{\"title\":\"new\",\"description\":\"new\",\"completed\":true}'

# Удалить
curl.exe -X DELETE http://localhost:8080/todos/1