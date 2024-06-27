package main

import (
	"fmt"
)

func main() {
	// マップに新しいキーと値を追加
	m := map[string]int{
		"apple":  1,
		"banana": 2,
	}
	m["orange"] = 3
	fmt.Println(m) // map[apple:1 banana:2 orange:3]

	// マップの長さ（要素数）を返す
	length := len(m)
	fmt.Println(length) // 3

	// マップから要素を削除
	delete(m, "banana")
	fmt.Println(m) // map[apple:1 orange:3]

	// マップの要素を個別に操作
	m["apple"] = 10
	fmt.Println(m) // map[apple:10 orange:3]

	// マップを範囲指定でループ処理
	// Key: apple, Value: 10 ...
	for key, value := range m {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}

	// キーの存在チェック
	if value, exists := m["grape"]; exists {
		fmt.Println("Grape exists:", value)
	} else {
		fmt.Println("Grape does not exist")
	}
}
