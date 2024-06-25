package main

import "fmt"

// 定数の宣言（iota使用）
const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

// 定数の宣言
const PI = 3.14159

func main() {
	// 定数（iota）の例
	today := Wednesday
	fmt.Printf("今日の曜日番号は%vです。\n", today) // "今日の曜日番号は3です。"と出力される

	// リテラルの使用例
	name := "アリス"
	greeting := "こんにちは！" + name
	fmt.Println(greeting) // "こんにちは！アリス"と出力される

	// 定数とリテラルを組み合わせた使用例
	var radius float64 = 5.0
	area := PI * radius * radius
	fmt.Printf("半径 %.2f の円の面積は %.2f\n", radius, area)
	// "半径 5.00 の円の面積は 78.54"と出力される
}
