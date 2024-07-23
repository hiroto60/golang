package controller

import (
	"fmt"
	"golang/interface/model"
	"log"

	"gorm.io/gorm"
)

type TodoController3 struct {
	Model *model.TodoModel2
}

// model.TodoModel3に依存している=gorm.DBを知っている
func NewTodoController3(db *gorm.DB) *TodoController3 {
	m := model.NewTodoModel2(db)
	return &TodoController3{Model: m}
}

func (c *TodoController2) GetPosts3() ([]model.Todo, error) {
	posts, err := c.Model.GetTodos2()
	if err != nil {
		log.Println("Failed to get posts2: ", err)
		return nil, err
	}

	fmt.Println("2:", posts)

	return posts, nil
}
