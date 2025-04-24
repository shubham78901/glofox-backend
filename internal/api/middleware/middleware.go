package middleware

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		path := r.URL.Path
		method := r.Method

		next.ServeHTTP(w, r)

		latency := time.Since(start)
		log.Printf("%s %s %s", method, path, latency)
	})
}

func ErrorHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(Response{
					Success: false,
					Message: "Internal server error",
				})
			}
		}()

		next.ServeHTTP(w, r)
	})
}
