package main

import "fmt"

func main() {
	// Print関数
	fmt.Print("Hello, World!\n")

	name := "辻野かなみ"
	age := 25
	fmt.Printf("%sは%d歳です。\n", name, age)

	fmt.Print("名前を入力してください: ")
	fmt.Scan(&name)
	fmt.Print("年齢を入力してください: ")
	fmt.Scan(&age)
	fmt.Printf("%sは%d歳です。\n", name, age)

}
