package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/hiroto60/golang/model"
)

type MockPostModel struct {
	Posts []model.Post
	Err   error
}

func TestGetPosts(t *testing.T) {
	// テスト用のデータ
	testTodos := []model.Post{
		{ID: 1, Title: "title1", Content: "content1", Likes: 0, CreatedAt: time.Now().Truncate(time.Millisecond)},
		{ID: 2, Title: "title2", Content: "content2", Likes: 0, CreatedAt: time.Now().Truncate(time.Millisecond)},
	}

	mockModel := &MockPostModel{Posts: testTodos}

	// TodoController の初期化
	controller := NewBlogController(mockModel)

	// テスト用のリクエストとレスポンスの作成
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/posts", nil)

	// テスト対象のハンドラーを呼び出す
	controller.GetPosts(w, r)

	// レスポンスのステータスコードが 200 であることを検証
	if w.Code != http.StatusOK {
		t.Errorf("expected %v but got %v", http.StatusOK, w.Code)
	}

	// レスポンスのボディを取得
	var got []model.Post
	if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	// // レスポンスのボディが期待通りであることを検証
	if !reflect.DeepEqual(got, testTodos) {
		t.Errorf("expected %v but got %v", testTodos, got)
	}

}

func (m *MockPostModel) GetPosts(w http.ResponseWriter, r *http.Request) ([]model.Post, error) {
	return m.Posts, m.Err
}

func (m *MockPostModel) GetPost(w http.ResponseWriter, r *http.Request, id int) (*model.Post, error) {
	return nil, nil
}

func (m *MockPostModel) CreatePost(w http.ResponseWriter, r *http.Request, post *model.Post) error {
	return nil
}

func (m *MockPostModel) Likes(w http.ResponseWriter, r *http.Request, id int) error {
	return nil
}

func (m *MockPostModel) UpdatePost(w http.ResponseWriter, r *http.Request, id int, newPost *model.Post) error {
	return nil
}

func (m *MockPostModel) DeletePost(w http.ResponseWriter, r *http.Request, id int) error {
	return nil
}
