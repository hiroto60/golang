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

// BlogModelerというinterfaceeに依存しているが、具体的な実装には依存していない。渡すDBなどを気にしないで良い。
func NewBlogController(m model.BlogModeler) *BlogController {
	return &BlogController{Model: m}
}

func (c *BlogController) GetPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := c.Model.GetPosts(w, r)
	if err != nil {
		log.Println("Failed to get posts: ", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}

func (c *BlogController) GetPost(w http.ResponseWriter, r *http.Request, id int) {

	post, err := c.Model.GetPost(w, r, id)
	if err != nil {
		log.Println("Failed to get post: ", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(post)
}

func (c *BlogController) CreatePost(w http.ResponseWriter, r *http.Request) {
	var post model.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		log.Println("Failed to decode post: ", err)
		http.Error(w, "Failed to decode post", http.StatusBadRequest)
		return
	}

	err = c.Model.CreatePost(w, r, &post)
	if err != nil {
		log.Println("Failed to create post: ", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(post)

}

func (c *BlogController) Likes(w http.ResponseWriter, r *http.Request, id int) {
	err := c.Model.Likes(w, r, id)
	if err != nil {
		log.Println("Failed to update post: ", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (c *BlogController) UpdatePost(w http.ResponseWriter, r *http.Request, id int) {
	var post model.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		log.Println("Failed to decode post: ", err)
		http.Error(w, "Failed to decode post", http.StatusBadRequest)
		return
	}

	err = c.Model.UpdatePost(w, r, id, &post)
	if err != nil {
		log.Println("Failed to update post: ", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (c *BlogController) DeletePost(w http.ResponseWriter, r *http.Request, id int) {
	err := c.Model.DeletePost(w, r, id)
	if err != nil {
		log.Println("Failed to delete post: ", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
