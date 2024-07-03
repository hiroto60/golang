package main

import (
	"fmt"
	"time"
)

// goroutineとして実行される関数
func printGoroutine() {
	for i := 0; i < 20; i++ {
		fmt.Println("Goroutine:", i)
		time.Sleep(time.Second) // 1秒間スリープ
	}
}

func main() {
	go printGoroutine() // goroutineを開始

	for i := 0; i < 5; i++ {
		fmt.Println("Main:", i)
		time.Sleep(time.Second) // 1秒間スリープ
	}
}
