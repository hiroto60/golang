package controller

import (
	"fmt"
	"golang/interface/model"
	"log"
)

type TodoController struct {
	Model model.TodoModeler
}

// model.TodoModelerに依存している=gorm.DBを知らず、model.TodoModelerを実装していれば引数に取れる
func NewTodoController(m model.TodoModeler) *TodoController {
	return &TodoController{Model: m}
}

func (c *TodoController) GetPosts() ([]model.Todo, error) {
	posts, err := c.Model.GetTodos()
	if err != nil {
		log.Println("Failed to get posts: ", err)
		return nil, err
	}

	fmt.Println(posts)

	return posts, nil
}
