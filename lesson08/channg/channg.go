package main

import (
	"fmt"
	"time"
)

// 与えられた数値の配列の合計を計算し、結果をポインタ経由で返す関数
func sumNumbers(numbers []int, result *int) {
	sum := 0
	for _, num := range numbers {
		sum += num
		fmt.Println("sumNumbers:", num)
		time.Sleep(time.Microsecond) // わずかな遅延を追加
	}
	*result = sum // 結果をポインタ経由で設定
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var sum1, sum2 int

	// 配列を2つに分割し、それぞれの部分配列に対してgoroutineを起動
	go sumNumbers(numbers[:len(numbers)/2], &sum1)
	go sumNumbers(numbers[len(numbers)/2:], &sum2)

	// goroutineの実行が完了するのを待たずに、sum1とsum2の値を出力
	fmt.Println("Sum 1:", sum1)
	fmt.Println("Sum 2:", sum2)
	fmt.Println("Total sum:", sum1+sum2)
}
