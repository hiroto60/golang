package model

import (
	"net/http"
	"time"

	"gorm.io/gorm"
)

type BlogModeler interface {
	GetPosts(w http.ResponseWriter, r *http.Request) (post []Post, err error)
	GetPost(w http.ResponseWriter, r *http.Request, id int) (*Post, error)
	CreatePost(w http.ResponseWriter, r *http.Request, post *Post) error
	DeletePost(w http.ResponseWriter, r *http.Request, id int) error
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

func (m *BlogModel) GetPosts(w http.ResponseWriter, r *http.Request) ([]Post, error) {
	var posts []Post
	result := m.DB.Find(&posts)
	if result.Error != nil {
		http.Error(w, "Failed to get posts", http.StatusInternalServerError)
		return nil, result.Error
	}

	return posts, nil

}

func (m *BlogModel) GetPost(w http.ResponseWriter, r *http.Request, id int) (*Post, error) {
	var post Post
	result := m.DB.First(&post, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			http.Error(w, "Post not found", http.StatusBadRequest)
			return nil, result.Error
		}
		http.Error(w, "Failed to get post", http.StatusInternalServerError)
		return nil, result.Error
	}

	return &post, nil
}

func (m *BlogModel) CreatePost(w http.ResponseWriter, r *http.Request, post *Post) error {

	result := m.DB.Create(&post)
	if result.Error != nil {
		http.Error(w, "Failed to create post", http.StatusInternalServerError)
		return result.Error
	}

	return nil
}

func (m *BlogModel) DeletePost(w http.ResponseWriter, r *http.Request, id int) error {
	var post Post

	result := m.DB.Delete(&post, id)
	if result.Error != nil {
		http.Error(w, "Failed to delete post", http.StatusInternalServerError)
		return result.Error
	}

	return nil
}
