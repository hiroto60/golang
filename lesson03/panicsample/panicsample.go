package main

import "fmt"

func processData(data string) {
	if len(data) == 0 {
		panic("データが空です")
	}
	fmt.Println("データの長さ:", len(data))
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("パニックが発生しました:", r)
		}
	}()

	processData("こんにちは")
	processData("")
	fmt.Println("プログラムが終了しました")
}
