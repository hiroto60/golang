package main

import "fmt"

func main() {
	score := 60 // 成績

	if score == 100 {
		fmt.Println("満点です！")
	} else if score >= 80 {
		fmt.Println("よくできました！")
	} else if score >= 60 {
		fmt.Println("合格です！")
	} else {
		fmt.Println("赤点です。。。")
	}
}
