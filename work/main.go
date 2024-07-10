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

	http.HandleFunc("/posts", server.HandlePosts)
	http.HandleFunc("/posts/", server.HandlePost)

	log.Fatal(s.ListenAndServe())
}
