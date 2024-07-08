package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/rs/cors"
)

// Book 構造体を定義
type Book struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title"`
	Price     int       `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

var db *gorm.DB

func main() {
	// データベースに接続
	dsn := "root:_MySQLPassword1@tcp(localhost:3306)/bookstore?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("データベースへの接続に失敗しました:", err)
	}

	// マイグレーションを実行
	db.AutoMigrate(&Book{})

	// ルートハンドラの登録
	http.HandleFunc("/books", handleBooks)
	http.HandleFunc("/books/", handleBook)

	// CORSの設定
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://127.0.0.1:5500"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type"},
	})

	// ハンドラにCORSミドルウェアを適用
	handler := c.Handler(http.DefaultServeMux)

	// サーバーの起動
	log.Fatal(http.ListenAndServe(":8080", handler))
}

// handleBooks は /books エンドポイントのリクエストを処理
func handleBooks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// 全ての本を取得
		getBooks(w, r)
	case http.MethodPost:
		// 新しい本を作成
		createBook(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// handleBook は /books/:id エンドポイントのリクエストを処理
func handleBook(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Path[len("/books/"):])
	if err != nil {
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case http.MethodGet:
		// 特定の本を取得
		getBook(w, r, id)
	case http.MethodPut:
		// 特定の本を更新
		updateBook(w, r, id)
	case http.MethodDelete:
		// 特定の本を削除
		deleteBook(w, r, id)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// getBooks は全ての本を取得
func getBooks(w http.ResponseWriter, r *http.Request) {
	var books []Book
	result := db.Find(&books)
	if result.Error != nil {
		http.Error(w, "クエリの実行に失敗しました", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(books)
}

// createBook は新しい本を作成
func createBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	// 未実装
	// json.NewDecoder(r.Body).Decode(&book)
	// book.ID = len(books) + 1
	// book.CreatedAt = time.Now()
	// books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

// getBook は特定の本を取得
func getBook(w http.ResponseWriter, r *http.Request, id int) {
	// 未実装
	// for _, book := range books {
	//   if book.ID == id {
	//     json.NewEncoder(w).Encode(book)
	//     return
	//   }
	// }
	http.NotFound(w, r)
}

// updateBook は特定の本を更新
func updateBook(w http.ResponseWriter, r *http.Request, id int) {
	var book Book
	result := db.First(&book, id)
	if result.Error != nil {
		http.NotFound(w, r)
		return
	}

	var updatedBook Book
	json.NewDecoder(r.Body).Decode(&updatedBook)
	book.Title = updatedBook.Title
	book.Price = updatedBook.Price
	result = db.Save(&book)
	if result.Error != nil {
		http.Error(w, "本の更新に失敗しました", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(book)
}

// deleteBook は特定の本を削除
func deleteBook(w http.ResponseWriter, r *http.Request, id int) {
	// 未実装
	// for i, book := range books {
	//   if book.ID == id {
	//     books = append(books[:i], books[i+1:]...)
	//     w.WriteHeader(http.StatusNoContent)
	//     return
	//   }
	// }
	http.NotFound(w, r)
}
