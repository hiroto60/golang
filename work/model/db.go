package model

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	//.envから読み取り
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	dbName := os.Getenv("MYSQL_DATABASE")
	port := os.Getenv("MYSQL_PORT")

	dsn := fmt.Sprintf("%s:%s@tcp(localhost:%s)/%s", user, password, port, dbName)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("データベースへの接続に失敗しました:", err)
		return
	}
}
