package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/hiroto60/golang/model"
)

// BlogModeler interfaceを実装する構造体
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

	controller := NewBlogController(mockModel)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/posts", nil)

	controller.GetPosts(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("expected %v but got %v", http.StatusOK, w.Code)
	}

	// レスポンスのボディを取得
	var got []model.Post
	t.Errorf("%v", w.Body)
	if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
		t.Errorf("failed to decode response: %v", err)
	}

	// レスポンスのボディが期待通りであることを検証
	if !reflect.DeepEqual(got, testTodos) {
		t.Errorf("expected %v but got %v", testTodos, got)
	}

}

func TestGetPost(t *testing.T) {
	testPost := model.Post{
		ID:        1,
		Title:     "title1",
		Content:   "content1",
		Likes:     0,
		CreatedAt: time.Now().Truncate(time.Millisecond),
	}

	mockModel := &MockPostModel{
		Posts: []model.Post{testPost},
	}

	controller := NewBlogController(mockModel)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/posts/1", nil)
	id := 1

	controller.GetPost(w, r, id)

	if w.Code != http.StatusOK {
		t.Errorf("expected %v but got %v", http.StatusOK, w.Code)
	}

	// レスポンスのボディを取得
	var got model.Post
	if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
		t.Errorf("failed to decode response: %v", err)
	}

	// レスポンスのボディが期待通りであることを検証
	if !reflect.DeepEqual(got, testPost) {
		t.Errorf("expected %v but got %v", testPost, got)
	}
}

func TestCreatePost(t *testing.T) {

	testPost := model.Post{
		ID:        1,
		Title:     "title1",
		Content:   "content1",
		Likes:     0,
		CreatedAt: time.Now().Truncate(time.Millisecond),
	}

	mockModel := &MockPostModel{}
	controller := NewBlogController(mockModel)

	postBody, err := json.Marshal(testPost)
	if err != nil {
		t.Fatalf("failed to marshal post: %v", err)
	}

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/posts", bytes.NewBuffer(postBody))
	r.Header.Set("Content-Type", "application/json")

	controller.CreatePost(w, r)

	if w.Code != http.StatusCreated {
		t.Errorf("expected %v but got %v", http.StatusCreated, w.Code)
	}
}

func TestLikes(t *testing.T) {
	// テスト用のデータ
	testPost := model.Post{
		ID:      1,
		Title:   "title1",
		Content: "content1",
		Likes:   0,
	}

	mockModel := &MockPostModel{
		Posts: []model.Post{testPost},
	}

	controller := NewBlogController(mockModel)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/posts/1/likes", nil)

	id := 1

	controller.Likes(w, r, id)

	if w.Code != http.StatusOK {
		t.Errorf("expected %v but got %v", http.StatusOK, w.Code)
	}
}

func TestUpdatePost(t *testing.T) {

	initialPost := model.Post{
		ID:        1,
		Title:     "title",
		Content:   "content",
		Likes:     0,
		CreatedAt: time.Now().Truncate(time.Millisecond),
	}

	updatedPost := model.Post{
		ID:        1,
		Title:     "updated title",
		Content:   "updated content",
		Likes:     0,
		CreatedAt: time.Now().Truncate(time.Millisecond),
	}

	mockModel := &MockPostModel{
		Posts: []model.Post{initialPost},
	}

	controller := NewBlogController(mockModel)

	postBody, err := json.Marshal(updatedPost)
	if err != nil {
		t.Fatalf("failed to marshal post: %v", err)
	}

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPut, "/posts/1", bytes.NewBuffer(postBody))
	r.Header.Set("Content-Type", "application/json")

	id := 1

	controller.UpdatePost(w, r, id)

	if w.Code != http.StatusOK {
		t.Errorf("expected %v but got %v", http.StatusOK, w.Code)
	}
}

func TestDeletePost(t *testing.T) {

	testPost := model.Post{
		ID:      1,
		Title:   "title1",
		Content: "content1",
		Likes:   0,
	}

	mockModel := &MockPostModel{
		Posts: []model.Post{testPost},
	}

	controller := NewBlogController(mockModel)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodDelete, "/posts/1", nil)

	id := 1

	controller.DeletePost(w, r, id)

	if w.Code != http.StatusNoContent {
		t.Errorf("expected %v but got %v", http.StatusNoContent, w.Code)
	}
}

func (m *MockPostModel) GetPosts(w http.ResponseWriter, r *http.Request) ([]model.Post, error) {
	return m.Posts, m.Err
}

func (m *MockPostModel) GetPost(w http.ResponseWriter, r *http.Request, id int) (*model.Post, error) {
	for _, post := range m.Posts {
		if post.ID == id {
			return &post, nil
		}
	}
	return nil, m.Err
}

func (m *MockPostModel) CreatePost(w http.ResponseWriter, r *http.Request, post *model.Post) error {
	if m.Err != nil {
		return m.Err
	}

	return nil
}

func (m *MockPostModel) Likes(w http.ResponseWriter, r *http.Request, id int) error {
	if m.Err != nil {
		return m.Err
	}

	return nil
}

func (m *MockPostModel) UpdatePost(w http.ResponseWriter, r *http.Request, id int, newPost *model.Post) error {
	if m.Err != nil {
		return m.Err
	}

	return nil
}

func (m *MockPostModel) DeletePost(w http.ResponseWriter, r *http.Request, id int) error {
	if m.Err != nil {
		return m.Err
	}

	return nil
}
