package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/hiroto60/golang/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// モックの準備
func setupMockDB() (*gorm.DB, sqlmock.Sqlmock, error) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	dialector := mysql.New(mysql.Config{
		DriverName:                "mysql",
		DSN:                       "sqlmock_db_0",
		Conn:                      mockDB,
		SkipInitializeWithVersion: true,
	})

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}

	return db, mock, nil
}

func TestGetPosts2(t *testing.T) {
	// モックデータベースのセットアップ
	db, mock, err := setupMockDB()
	if err != nil {
		t.Fatalf("failed to setup mock database: %v", err)
	}

	// モックレスポンスの準備
	mockPosts := []model.Post{
		{ID: 1, Title: "title1", Content: "content1", Likes: 0},
		{ID: 2, Title: "title2", Content: "content2", Likes: 0},
	}

	rows := sqlmock.NewRows([]string{"id", "title", "content", "likes"}).
		AddRow(mockPosts[0].ID, mockPosts[0].Title, mockPosts[0].Content, mockPosts[0].Likes).
		AddRow(mockPosts[1].ID, mockPosts[1].Title, mockPosts[1].Content, mockPosts[1].Likes)

	mock.ExpectQuery("^SELECT \\* FROM `posts`$").WillReturnRows(rows)

	// コントローラーの初期化
	controller := NewBlogController2(db)

	// テスト用のリクエストとレスポンスの作成
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/posts", nil)

	// テスト対象のハンドラーを呼び出す
	posts, err := controller.GetPosts(w, r)
	if err != nil {
		t.Fatalf("expected no error, but got %v", err)
	}

	// レスポンスのステータスコードが 200 であることを検証
	if w.Code != http.StatusOK {
		t.Errorf("expected %v but got %v", http.StatusOK, w.Code)
	}

	// レスポンスのボディが期待通りであることを検証
	if len(posts) != len(mockPosts) {
		t.Errorf("expected %v posts but got %v", len(mockPosts), len(posts))
	}

	for i, post := range posts {
		if post.ID != mockPosts[i].ID || post.Title != mockPosts[i].Title || post.Content != mockPosts[i].Content || post.Likes != mockPosts[i].Likes {
			t.Errorf("expected post %v but got %v", mockPosts[i], post)
		}
	}

	// モックの期待を満たしていることを検証
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %v", err)
	}
}
