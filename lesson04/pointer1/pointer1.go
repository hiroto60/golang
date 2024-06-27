package main

import (
	"fmt"
)

func main() {
	var x int = 10
	var p *int = &x // xのアドレスをポインターpに代入
	fmt.Println(p)  // ポインターの値を出力（0xc00000a0b8など）
	fmt.Println(*p) // ポインターpが指す値を出力（10）
	*p = 20         // ポインターpが指す値を変更
	fmt.Println(x)  // xの値が変更されている（20）
	fmt.Println(p)  // ポインターの値は変わらない
}
