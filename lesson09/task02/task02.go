package main

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Book struct {
	ID        int
	Title     string
	Price     int
	CreatedAt time.Time
}

func main() {
	dsn := "root:_MySQLPassword1@tcp(localhost:3306)/bookstore?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("データベースへの接続に失敗しました:", err)
	}

	fmt.Printf("検索するタイトルを入力してください:")

	var title string
	_, err = fmt.Scan(&title)
	if err != nil {
		fmt.Println("入力エラーです。")
		return
	}

	var book Book
	result := db.First(&book, "title = ?", title)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			fmt.Println("指定されたタイトルのレコードが見つかりませんでした")
		} else {
			fmt.Println("クエリの実行に失敗しました:", result.Error)
		}
		return
	}

	// 結果を出力
	fmt.Println(book.ID, book.Title, book.Price, book.CreatedAt)

}
