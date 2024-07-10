package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/hiroto60/golang/model"
)

type BlogController struct {
	Model model.BlogModeler
}

func NewBlogController(m model.BlogModeler) *BlogController {
	return &BlogController{Model: m}
}

func GetPosts() {
	// Get post
}

func (c *BlogController) CreatePost(w http.ResponseWriter, r *http.Request) {
	var post model.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		log.Println("Failed to decode post: ", err)
		http.Error(w, "Failed to decode post", http.StatusBadRequest)
	}

	err = c.Model.CreatePost(w, r, &post)
	if err != nil {
		log.Println("Failed to create post: ", err)
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(post)

}
