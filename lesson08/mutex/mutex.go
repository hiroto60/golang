package main

import (
	"fmt"
	"sync"
)

// 同期機能を備えたカウンターを表す構造体
type Counter struct {
	mu    sync.Mutex
	value int
}

// カウンターの値を1つインクリメントする
func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// カウンターの現在の値を返す
func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

// カウンターをインクリメントするゴルーチンの関数
func incrementCounter(counter *Counter, wg *sync.WaitGroup) {
	defer wg.Done()
	counter.Increment()
}

func main() {
	var counter Counter   // カウンター変数を初期化
	var wg sync.WaitGroup // WaitGroup を初期化

	// 1000個のゴルーチンを作成し、カウンターをインクリメント
	for i := 0; i < 1000; i++ {
		wg.Add(1)                          // WaitGroup のカウンターを1つ増やす
		go incrementCounter(&counter, &wg) // ゴルーチンを作成し、カウンターをインクリメント
	}

	wg.Wait() // すべてのゴルーチンが完了するまで待機

	// 最終的なカウンターの値を出力
	fmt.Println("カウンター:", counter.Value())
}
