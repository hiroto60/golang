package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("[%s] %s %s %s", time.Now().Format(time.RFC3339), r.Method, r.URL.Path, time.Since(start))
	})
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	http.Handle("/hello", loggingMiddleware(http.HandlerFunc(hello)))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
