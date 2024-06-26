package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("処理を終了します")

	var a, b int

	fmt.Print("割られる数を入力してください：")
	_, err := fmt.Scanln(&a)
	if err != nil {
		fmt.Println("エラー:数値を入力してください")
		return
	}

	fmt.Print("割る数を入力してください：")
	_, err = fmt.Scanln(&b)
	if err != nil {
		fmt.Println("エラー:数値を入力してください")
		return
	}

	if b == 0 {
		fmt.Println("エラー:0で割り算しないでください")
		return
	}

	c := float64(a) / float64(b)

	fmt.Printf("%d ÷ %d = %.1f\n", a, b, c)

}
