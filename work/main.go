package main

import (
	"log"
	"net/http"

	"github.com/hiroto60/golang/server"
	"github.com/rs/cors"
)

func main() {
	// CORSの設定
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://127.0.0.1:5500"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type"},
	})

	// ハンドラにCORSミドルウェアを適用
	handler := c.Handler(http.DefaultServeMux)
	s := http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	http.HandleFunc("/posts", server.HandlePosts)
	http.HandleFunc("/posts/", server.HandlePost)

	log.Fatal(s.ListenAndServe())
}
