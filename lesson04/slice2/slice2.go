package main

import (
	"fmt"
	"sort"
)

func main() {
	// スライスの末尾に新しい要素を追加
	numbers := []int{1, 2, 3}
	numbers = append(numbers, 4, 5)
	fmt.Println(numbers) // [1 2 3 4 5]

	// スライスの長さ（要素数）を返す
	length := len(numbers)
	fmt.Println(length) // 5

	// スライスの容量を返す
	capacity := cap(numbers)
	fmt.Println(capacity) // 6

	// スライスの一部を切り出す
	slice := numbers[1:4]
	fmt.Println(slice) // [2 3 4]

	// スライスのコピー
	src := []int{1, 2, 3}
	dst := make([]int, len(src))
	copy(dst, src)
	fmt.Println(dst) // [1 2 3]
	// 元のスライスを変更してもコピー先には影響しない
	src[0] = 10
	fmt.Println(src) // [10 2 3]
	fmt.Println(dst) // [1 2 3]

	// スライスの要素を個別に操作
	numbers[0] = 10
	fmt.Println(numbers) // [10 2 3 4 5]

	// スライスを範囲指定でループ処理
	// Index: 0, Value: 10 ...
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// スライスのソート
	randNumbers := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3}
	sort.Ints(randNumbers)
	fmt.Println(randNumbers) // [1 1 2 3 3 4 5 5 6 9]
}
