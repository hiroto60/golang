package main

import "fmt"

func main() {
	score := 72 // 成績

	switch {
	case score == 100:
		fmt.Println("満点です！")
	case score >= 80:
		fmt.Println("よくできました！")
	case score >= 60:
		fmt.Println("合格です！")
	default:
		fmt.Println("赤点です。。。")
	}
}
