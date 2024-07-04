package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Book struct {
	ID        int
	Title     string
	Price     int
	CreatedAt time.Time
}

func main() {
	// データベースに接続
	db, err := sql.Open("mysql", "root:_MySQLPassword1@tcp(localhost:3306)/bookstore?charset=utf8mb4&parseTime=true&loc=Local")
	if err != nil {
		fmt.Println("データベースへの接続に失敗しました:", err)
		return
	}
	defer db.Close()

	// トランザクションを開始
	tx, err := db.Begin()
	if err != nil {
		fmt.Println("トランザクションの開始に失敗しました:", err)
		return
	}

	// プリペアドステートメントを使用してデータを挿入
	insertStmt, err := tx.Prepare("INSERT INTO books(title, price, created_at) VALUES(?, ?, ?)")
	if err != nil {
		fmt.Println("プリペアドステートメントの作成に失敗しました:", err)
		tx.Rollback()
		return
	}
	defer insertStmt.Close()

	// 複数のデータを挿入
	books := []Book{
		{Title: "Book1", Price: 1000, CreatedAt: time.Now()},
		{Title: "Book2", Price: 2000, CreatedAt: time.Now()},
		{Title: "Book3", Price: 3000, CreatedAt: time.Now()},
	}

	for _, book := range books {
		_, err := insertStmt.Exec(book.Title, book.Price, book.CreatedAt)
		if err != nil {
			fmt.Println("データの挿入に失敗しました:", err)
			tx.Rollback()
			return
		}
	}

	// トランザクションをコミット
	err = tx.Commit()
	if err != nil {
		fmt.Println("トランザクションのコミットに失敗しました:", err)
		tx.Rollback()
		return
	}

	fmt.Println("データの挿入が成功しました")
}
