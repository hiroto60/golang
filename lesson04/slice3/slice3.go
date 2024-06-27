package main

import "fmt"

func main() {
	// 元のスライス
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("元のスライス:", numbers)

	// 既存のスライスから新しいスライスを作成
	newSlice := numbers[2:]
	fmt.Println("新しいスライス:", newSlice)

	// 新しいスライスの要素を変更
	newSlice[0] = 30
	fmt.Println("変更後の新しいスライス:", newSlice)

	// 元のスライスの要素が変更されていることを確認
	fmt.Println("変更後の元のスライス:", numbers)
}
