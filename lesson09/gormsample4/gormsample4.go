package main

import (
	"fmt"
	"time"

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
	// データベースに接続
	dsn := "root:_MySQLPassword1@tcp(localhost:3306)/bookstore?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("データベースへの接続に失敗しました:", err)
		return
	}

	// トランザクションを開始
	tx := db.Begin()

	// 複数のデータを挿入
	books := []Book{
		{Title: "Book1", Price: 1000, CreatedAt: time.Now()},
		{Title: "Book2", Price: 2000, CreatedAt: time.Now()},
		{Title: "Book3", Price: 3000, CreatedAt: time.Now()},
	}

	for _, book := range books {
		if err := tx.Create(&book).Error; err != nil {
			fmt.Println("データの挿入に失敗しました:", err)
			tx.Rollback()
			return
		}
	}

	// トランザクションをコミット
	if err := tx.Commit().Error; err != nil {
		fmt.Println("トランザクションのコミットに失敗しました:", err)
		return
	}

	fmt.Println("データの挿入が成功しました")
}
