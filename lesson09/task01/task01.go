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

	selectStmt, err := db.Prepare("SELECT * FROM books WHERE title Like ?")
	if err != nil {
		fmt.Println("Prepare error:", err)
		return
	}

	var selectedBooks []Book

	rows, err := selectStmt.Query("%" + title + "%")
	if err != nil {
		fmt.Println("検索に失敗しました:", err)
		return
	}

	for rows.Next() {
		var book Book
		err = rows.Scan(&book.ID, &book.Title, &book.Price, &book.CreatedAt)
		if err != nil {
			fmt.Println("検索結果の読み込みに失敗しました:", err)
			return
		}

		selectedBooks = append(selectedBooks, book)
	}

	if len(selectedBooks) == 0 {
		fmt.Println("指定されたタイトルのレコードが見つかりませんでした")
		return
	} else {
		for _, book := range selectedBooks {
			fmt.Printf("%v %v %v %v\n", book.ID, book.Title, book.Price, book.CreatedAt)
		}
	}

}
