package router

import (
	"log"
	"net/http"
	"time"
	"todos_manager/internal/handlers"
)

func NewRouter(handler *handlers.TodoHandler) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /todos", handler.CreateTodo)
	mux.HandleFunc("GET /todos", handler.ListTodos)
	mux.HandleFunc("GET /todos/{id}", handler.GetTodo)
	mux.HandleFunc("PUT /todos/{id}", handler.UpdateTodo)
	mux.HandleFunc("DELETE /todos/{id}", handler.DeleteTodo)
	mux.HandleFunc("PATCH /todos/{id}/complete", handler.CompleteTodo)
	return logMiddleware(mux)
}

func logMiddleware(mux http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		lrw := &loggingResponseWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}
		mux.ServeHTTP(lrw, r)
		duration := time.Since(start)
		log.Printf("[%s] %s %s - %d - %v",
			time.Now().Format("2006-01-02 15:04:05"),
			r.Method,
			r.URL.Path,
			lrw.statusCode,
			duration,
		)
	})
}

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}
