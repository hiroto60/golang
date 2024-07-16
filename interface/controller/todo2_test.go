package controller

import (
	"fmt"
	"golang/interface/model"
	"testing"
)

// モックモデルの定義
type MockTodoModel2 struct {
	model.TodoModel2
	MockGetTodos2 func() ([]model.Todo, error)
}

func (m *MockTodoModel2) GetTodos2() ([]model.Todo, error) {
	return m.MockGetTodos2()
}

// テストコード
func TestTodoController2_GetPosts2_Success(t *testing.T) {
	// モックのセットアップ
	mockModel := &MockTodoModel2{
		MockGetTodos2: func() ([]model.Todo, error) {
			return []model.Todo{
				{ID: 1, Title: "Task1"},
				{ID: 2, Title: "Task2"},
			}, nil
		},
	}

	// コントローラのセットアップ
	controller := NewTodoController2(&mockModel.TodoModel2)

	// メソッドの実行
	todos, err := controller.GetPosts2()
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
