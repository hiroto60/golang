package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

// Book 構造体を定義
type Book struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Price     int       `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

// books スライスを宣言
var books []Book

func main() {
	// サンプルデータの作成
	books = append(books, Book{ID: 1, Title: "Book 1", Price: 1000, CreatedAt: time.Now()})
	books = append(books, Book{ID: 2, Title: "Book 2", Price: 1500, CreatedAt: time.Now()})
	books = append(books, Book{ID: 3, Title: "Book 3", Price: 2000, CreatedAt: time.Now()})

	// ルートハンドラの登録
	http.HandleFunc("/books", handleBooks)
	http.HandleFunc("/books/", handleBook)

	// サーバーの起動
	log.Fatal(http.ListenAndServe(":8080", nil))
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
	json.NewEncoder(w).Encode(books)
}

// createBook は新しい本を作成
func createBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	json.NewDecoder(r.Body).Decode(&book)
	book.ID = len(books) + 1
	book.CreatedAt = time.Now()
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

// getBook は特定の本を取得
func getBook(w http.ResponseWriter, r *http.Request, id int) {
	for _, book := range books {
		if book.ID == id {
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	http.NotFound(w, r)
}

// updateBook は特定の本を更新
func updateBook(w http.ResponseWriter, r *http.Request, id int) {
	var updatedBook Book
	json.NewDecoder(r.Body).Decode(&updatedBook)

	for i, book := range books {
		if book.ID == id {
			updatedBook.ID = book.ID
			updatedBook.CreatedAt = book.CreatedAt
			books[i] = updatedBook
			json.NewEncoder(w).Encode(updatedBook)
			return
		}
	}
	http.NotFound(w, r)
}

// deleteBook は特定の本を削除
func deleteBook(w http.ResponseWriter, r *http.Request, id int) {
	for i, book := range books {
		if book.ID == id {
			books = append(books[:i], books[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.NotFound(w, r)
}
