package server

import (
	"net/http"

	"github.com/hiroto60/golang/controller"
	"github.com/hiroto60/golang/model"
)

func HandlePosts2(w http.ResponseWriter, r *http.Request) {
	db := model.GetDB()
	bc := controller.NewBlogController2(db)

	switch r.Method {
	case http.MethodGet:
		bc.GetPosts(w, r)
	case http.MethodPost:
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
