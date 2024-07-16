package model

import "gorm.io/gorm"

type Todo2 struct {
	ID    int
	Title string
}

type TodoModel2 struct {
	DB *gorm.DB
}

func NewTodoModel2(db *gorm.DB) *TodoModel2 {
	return &TodoModel2{DB: db}
}

func (m *TodoModel2) GetTodos2() ([]Todo, error) {
	var todos []Todo
	if err := m.DB.Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}
