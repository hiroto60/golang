package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int = 10
	var f64 float64 = float64(i) // int型からfloat64型への変換
	fmt.Printf("型:%T, 値:%v\n", f64, f64)

	var f float64 = 3.14
	i = int(f) // float64型からint型への変換
	fmt.Printf("型:%T, 値:%v\n", i, i)

	var i64 int64 = int64(i) // int型からint64型への変換も、明示的な型変換が必要
	fmt.Printf("型:%T, 値:%v\n", i64, i64)
	var i32 int32 = int32(i) // int型からint32型への変換も、明示的な型変換が必要
	fmt.Printf("型:%T, 値:%v\n", i32, i32)

	var str string = "golang"

	i, err := strconv.Atoi(str) // string型からint型への変換
	if err != nil {
		// エラー処理
		fmt.Println("文字列を整数に変換できませんでした:", err)
	}

	str = strconv.Itoa(i) // int型からstring型への変換
	fmt.Printf("型:%T, 値:%v\n", str, str)
}
