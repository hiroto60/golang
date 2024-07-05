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
	db, err := sql.Open("mysql", "root:_MySQLPassword1@tcp(localhost:3306)/bookstore?charset=utf8mb4&parseTime=true&loc=Local")
	if err != nil {
		fmt.Println("データベースへの接続に失敗しました:", err)
		return
	}
	defer db.Close()

	fmt.Printf("検索するタイトルを入力してください:")
	var title string
	_, err = fmt.Scan(&title)
	if err != nil {
		fmt.Println("入力エラーです。")
		return
	}

	selectStmt, err := db.Prepare("SELECT * FROM books WHERE title = ?")
	if err != nil {
		fmt.Println("Prepare error:", err)
		return
	}

	var selectedBook Book

	err = selectStmt.QueryRow(title).Scan(&selectedBook.ID, &selectedBook.Title, &selectedBook.Price, &selectedBook.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("該当する書籍が見つかりませんでした。")
		} else {
			fmt.Println("検索に失敗しました:", err)
		}
		return
	}

	fmt.Printf("%v %v %v %v\n", selectedBook.ID, selectedBook.Title, selectedBook.Price, selectedBook.CreatedAt)

}
