package main

import (
	"fmt"
	"time"
)

// ch1に値を送信する関数
func sendToCh1(ch chan<- int) {
	time.Sleep(2 * time.Second)
	ch <- 42
}

// ch2に値を送信する関数
func sendToCh2(ch chan<- string) {
	time.Sleep(1 * time.Second)
	ch <- "hello"
}

// ch3に値を送信する関数
func sendToCh3(ch chan<- bool) {
	time.Sleep(10 * time.Second)
	ch <- false
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan string)
	ch3 := make(chan bool)

	// それぞれのチャネルに値を送信するgoroutineを起動
	go sendToCh1(ch1)
	go sendToCh2(ch2)
	go sendToCh3(ch3)

	// 3回のループで、selectを使用してチャネルからの値の受信を待機
	for i := 0; i < 3; i++ {
		select {
		case value := <-ch1:
			fmt.Println("Received from ch1:", value)
		case value := <-ch2:
			fmt.Println("Received from ch2:", value)
		case value := <-ch3:
			fmt.Println("Received from ch3:", value)
		case <-time.After(3 * time.Second):
			fmt.Println("Timeout: No value received")
		}
	}
}
