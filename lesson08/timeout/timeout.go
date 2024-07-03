package main

import (
	"context"
	"fmt"
	"time"
)

// 長時間の処理を行う関数
func longRunningProcess(ctx context.Context, id int) {
	fmt.Printf("プロセス %d が開始しました\n", id)

	// コンテキストのキャンセルまたはタイムアウトを監視するためのselect文
	select {
	case <-time.After(5 * time.Second):
		fmt.Printf("プロセス %d が完了しました\n", id)
	case <-ctx.Done():
		fmt.Printf("プロセス %d がキャンセルされました\n", id)
	}
}

func main() {
	// タイムアウト付きのコンテキストを作成
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// 複数のゴルーチンを開始
	for i := 1; i <= 3; i++ {
		go longRunningProcess(ctx, i)
	}

	// コンテキストのタイムアウトまたはキャンセルを待機
	<-ctx.Done()

	// コンテキストがキャンセルされた理由を表示
	fmt.Println("コンテキストがキャンセルされました:", ctx.Err())
}
