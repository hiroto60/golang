package controller

import (
	"fmt"
	"golang/interface/model"
	"log"
)

// TodoModel2に依存している
type TodoController2 struct {
	Model *model.TodoModel2
}

func NewTodoController2(m *model.TodoModel2) *TodoController2 {
	return &TodoController2{Model: m}
}

func (c *TodoController2) GetPosts2() ([]model.Todo, error) {
	posts, err := c.Model.GetTodos2()
	if err != nil {
		log.Println("Failed to get posts2: ", err)
		return nil, err
	}

	fmt.Println("2:", posts)

	return posts, nil
}
