package server

import (
	"net/http"

	"github.com/hiroto60/golang/controller"
	"github.com/hiroto60/golang/model"
)

func HandlePosts(w http.ResponseWriter, r *http.Request) {
	db := model.GetDB()
	bm := model.NewBlogModel(db)
	bc := controller.NewBlogController(bm)

	switch r.Method {
	case http.MethodGet:
		bc.GetPosts(w, r)
	case http.MethodPost:
		bc.CreatePost(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func HandlePost(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
	case http.MethodPut:
	case http.MethodDelete:
	case http.MethodPost:
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
