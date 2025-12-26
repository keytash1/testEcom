FROM golang:1.21-alpine
WORKDIR /app
COPY . .
RUN go build -o todo-manager ./cmd/server
EXPOSE 8080
CMD ["./todo-manager"]