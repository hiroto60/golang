package main

import "fmt"

func main() {
	var scores []int
	var score int

	total := 0

	for {
		fmt.Println("点数を入力してください:")
		fmt.Scan(&score)
		if score == -1 {
			break
		}
		total += score
		scores = append(scores, score)
	}
	if len(scores) == 0 {
		fmt.Println("点数が入力されませんでした")

	} else {
		fmt.Printf("%v人のテストの平均点は%.1f点です。", len(scores), float64(total)/float64((len(scores))))
	}

}
