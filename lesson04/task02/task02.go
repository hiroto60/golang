package main

import "fmt"

func main() {
	var word string
	words := make([]string, 0)
	counts := make(map[rune]int)

	fmt.Println("英単語を入力してください（endと入力するとループを終了します）：")
	for {
		fmt.Scan(&word)

		if word == "end" {
			break
		}

		words = append(words, word)

		for _, r := range word {
			counts[r]++
		}
	}

	fmt.Println("入力された英単語：")
	for _, w := range words {
		fmt.Println(w)
	}

	fmt.Println("アルファベットごとの文字数：")
	for r, count := range counts {
		fmt.Printf("%c：%v\n", r, count)
	}

}
