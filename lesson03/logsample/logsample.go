package main

import (
	"log"
	"os"
)

func main() {
	// ログの出力先をファイルに変更
	file, err := os.Create("log.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	log.SetOutput(file)
	log.Println("ファイルにログを出力します")

	var count int = 10
	log.Printf("現在のカウント: %d\n", count)
}
