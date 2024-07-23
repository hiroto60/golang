package controller

import (
	"fmt"
	"golang/interface/model"
	"log"
)

type TodoController2 struct {
	Model *model.TodoModel2
}

// model.TodoModel2に依存している=gorm.DBを知っている
// NewTodoController2を呼び出すときにmodel.TodoModel2の構造体で定義された型を渡す必要がある(今回は*gorm.DB)=ormを変えた場合は、
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
