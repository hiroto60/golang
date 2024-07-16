package model

import (
	"net/http"

	"gorm.io/gorm"
)

type BlogModel2 struct {
	DB *gorm.DB
}

func NewBlogModel2(db *gorm.DB) *BlogModel2 {
	return &BlogModel2{DB: db}
}

func (m *BlogModel2) GetPosts(w http.ResponseWriter, r *http.Request) ([]Post, error) {
	var posts []Post
	result := m.DB.Find(&posts)
	if result.Error != nil {
		return nil, result.Error
	}

	return posts, nil
}
