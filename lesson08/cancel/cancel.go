package main

import (
	"context"
	"fmt"
	"time"
)

// 長時間の処理を行う関数
func longRunningProcess(ctx context.Context, id int) {
	fmt.Printf("プロセス %d が開始しました\n", id)

	// コンテキストのキャンセルを監視するためのselect文
	select {
	case <-time.After(5 * time.Second):
		fmt.Printf("プロセス %d が完了しました\n", id)
	case <-ctx.Done():
		fmt.Printf("プロセス %d がキャンセルされました\n", id)
	}
}

func main() {
	// キャンセル可能なコンテキストを作成
	ctx, cancel := context.WithCancel(context.Background())

	// 複数のゴルーチンを開始
	for i := 1; i <= 3; i++ {
		go longRunningProcess(ctx, i)
	}

	// 2秒後にコンテキストをキャンセル
	time.Sleep(2 * time.Second)
	cancel()

	// ゴルーチンの終了を待つために少し待機
	time.Sleep(1 * time.Second)
}
