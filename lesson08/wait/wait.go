package main

import (
	"fmt"
	"sync"
	"time"
)

// worker関数はゴルーチンとして実行される処理を含む
func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // ゴルーチンの完了を通知するためにDone()をdeferで呼び出す
	fmt.Printf("ゴルーチン %d が開始しました\n", id)
	time.Sleep(2 * time.Second) // 2秒間のスリープを行う
	fmt.Printf("ゴルーチン %d が終了しました\n", id)
}

func main() {
	var wg sync.WaitGroup // WaitGroupを作成

	// 5つのゴルーチンを作成して実行
	for i := 0; i < 5; i++ {
		wg.Add(1)         // WaitGroupのカウンタを1つ増やす
		go worker(i, &wg) // worker関数をゴルーチンとして実行
	}

	fmt.Println("ゴルーチンの完了を待っています...")
	wg.Wait() // すべてのゴルーチンの完了を待つ
	fmt.Println("すべてのゴルーチンが終了しました")
}
