package model

import "gorm.io/gorm"

type TodoModeler interface {
	GetTodos() ([]Todo, error)
}

type Todo struct {
	ID    int
	Title string
}

type TodoModel struct {
	DB *gorm.DB
}

func NewTodoModel(db *gorm.DB) TodoModeler {
	return &TodoModel{DB: db}
}

func (m *TodoModel) GetTodos() ([]Todo, error) {
	var todos []Todo
	m.DB.Find(&todos)
	return todos, nil

}
