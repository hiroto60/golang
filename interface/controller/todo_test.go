package controller

import (
	"fmt"
	"golang/interface/model"
	"testing"
)

// モックインターフェースの定義
type MockTodoModel struct {
	models []model.Todo
	err    error
}

func (m *MockTodoModel) GetTodos() ([]model.Todo, error) {
	return m.models, m.err
}

// テストコード
func TestTodoController(t *testing.T) {
	// モックのセットアップ
	mockModel := &MockTodoModel{
		models: []model.Todo{
			{ID: 1, Title: "Task1"},
			{ID: 2, Title: "Task2"},
		},
		err: nil,
	}

	// コントローラのセットアップ
	controller := NewTodoController(mockModel)

	// メソッドの実行
	todos, err := controller.GetPosts()
	if err != nil {
		t.Error("Expected no error, got ", err)
	}

	expected := []model.Todo{
		{ID: 1, Title: "Task1"},
		{ID: 2, Title: "Task2"},
	}
	if fmt.Sprint(todos) != fmt.Sprint(expected) {
		t.Errorf("Expected %v, got %v", expected, todos)
	}
}
