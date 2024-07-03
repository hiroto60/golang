package main

import (
	"context"
	"fmt"
)

// キーの型を定義
type keyType string

// ユーザーIDを表すキー
const userIDKey keyType = "userID"

// 認証済みかどうかを表すキー
const isAuthenticatedKey keyType = "isAuthenticated"

// 処理を行う関数
func processRequest(ctx context.Context) {
	// コンテキストからユーザーIDを取得
	userID := ctx.Value(userIDKey).(string)
	fmt.Println("ユーザーのリクエストを処理中:", userID)

	// コンテキストから認証状態を取得
	isAuthenticated := ctx.Value(isAuthenticatedKey).(bool)
	if isAuthenticated {
		fmt.Println("ユーザーは認証済みです")
	} else {
		fmt.Println("ユーザーは認証されていません")
	}
}

// ミドルウェアを表す関数
func middlewareAuth(next func(ctx context.Context), userID string, isAuthenticated bool) func(ctx context.Context) {
	return func(ctx context.Context) {
		// 新しいコンテキストにユーザーIDと認証状態を追加
		ctx = context.WithValue(ctx, userIDKey, userID)
		ctx = context.WithValue(ctx, isAuthenticatedKey, isAuthenticated)

		// 次の処理を呼び出す
		next(ctx)
	}
}

func main() {
	// コンテキストを作成
	ctx := context.Background()

	// ミドルウェアを適用して処理を実行
	middlewareAuth(processRequest, "user123", true)(ctx)
	middlewareAuth(processRequest, "user456", false)(ctx)
}
