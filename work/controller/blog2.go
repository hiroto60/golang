package controller

import (
	"net/http"

	"github.com/hiroto60/golang/model"
	"gorm.io/gorm"
)

type BlogController2 struct {
	Model model.BlogModel2
}

// NewBlogModelして直接BlogControllerに渡している=>controller層がmodel層に依存している。渡すDBを気にする必要がある。
func NewBlogController2(db *gorm.DB) *BlogController2 {
	m := model.NewBlogModel2(db)
	return &BlogController2{Model: *m}
}

func (c *BlogController2) GetPosts(w http.ResponseWriter, r *http.Request) ([]model.Post, error) {
	posts, err := c.Model.GetPosts(w, r)
	if err != nil {
		return nil, err
	}

	return posts, nil
}
