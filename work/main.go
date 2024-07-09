package main

import (
	"log"
	"net/http"

	"github.com/hiroto60/golang/server"
)

func main() {
	s := http.Server{
		Addr: ":8080",
	}

	http.HandleFunc("/books", server.HandlePosts)
	http.HandleFunc("/books/", server.HandlePost)

	log.Fatal(s.ListenAndServe())
}
