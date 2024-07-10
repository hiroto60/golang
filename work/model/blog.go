package model

import (
	"net/http"
	"time"

	"gorm.io/gorm"
)

type BlogModeler interface {
	GetPosts()
	CreatePost(w http.ResponseWriter, r *http.Request, post *Post) error
}

type Post struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Likes     int       `json:"likes"`
	CreatedAt time.Time `json:"created_at"`
}

type BlogModel struct {
	DB *gorm.DB
}

func NewBlogModel(DB *gorm.DB) BlogModeler {
	return &BlogModel{DB: DB}
}

func (m *BlogModel) GetPosts() {
	// Get post
}

func (m *BlogModel) CreatePost(w http.ResponseWriter, r *http.Request, post *Post) error {

	result := m.DB.Create(&post)
	if result.Error != nil {
		http.Error(w, "Failed to create post", http.StatusInternalServerError)
		return result.Error
	}

	return nil
}
