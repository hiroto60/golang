package main

import (
	"fmt"
	"time"
)

// 与えられた数値の配列の合計を計算し、結果をチャネルに送信する関数
func sumNumbers(numbers []int, ch chan<- int) {
	sum := 0
	for _, num := range numbers {
		sum += num
		fmt.Println("sumNumbers:", num)
		time.Sleep(time.Microsecond) // わずかな遅延を追加
	}
	ch <- sum // 計算結果をチャネルに送信
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ch := make(chan int) // 結果を受信するチャネルを作成

	// 配列を2つに分割し、それぞれの部分配列に対してgoroutineを起動
	go sumNumbers(numbers[:len(numbers)/2], ch)
	go sumNumbers(numbers[len(numbers)/2:], ch)

	// チャネルから結果を受信
	sum1, sum2 := <-ch, <-ch

	fmt.Println("Sum 1:", sum1)
	fmt.Println("Sum 2:", sum2)
	fmt.Println("Total sum:", sum1+sum2)
}
