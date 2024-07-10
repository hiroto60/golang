package server

import (
	"net/http"
	"strconv"

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
	db := model.GetDB()
	bm := model.NewBlogModel(db)
	bc := controller.NewBlogController(bm)

	//pathパラメータの取得
	id, err := strconv.Atoi(r.URL.Path[len("/books/"):])
	if err != nil {
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case http.MethodGet:
		bc.GetPost(w, r, id)
	case http.MethodPut:
		bc.UpdatePost(w, r, id)
	case http.MethodDelete:
		bc.DeletePost(w, r, id)
	case http.MethodPost:
		bc.Likes(w, r, id)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
