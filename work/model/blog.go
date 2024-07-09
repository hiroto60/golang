package model

import (
	"encoding/json"
	"net/http"
	"time"
)

type BlogModeler interface {
	GetPosts()
	CreatePost(w http.ResponseWriter, r *http.Request)
}

type Post struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Likes     int       `json:"likes"`
	CreatedAt time.Time `json:"created_at"`
}

func (p *Post) GetPosts() {
	// Get post
}

func (p *Post) CreatePost(w http.ResponseWriter, r *http.Request) {
	var post Post
	json.NewDecoder(r.Body).Decode(&post)

	result := DB.Create(&post)
	if result.Error != nil {
		http.Error(w, "Failed to create post", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(post)
}
