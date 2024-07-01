package main

import "fmt"

func main() {
	var anything interface{}
	anything = "Hello, World!"
	fmt.Println(anything) // "Hello, World!"

	anything = 42
	fmt.Println(anything) // 42

	anything = []int{1, 2, 3}
	fmt.Println(anything) // [1 2 3]

	// 型アサーションの例1: インターフェース型の値を文字列型にキャストする
	str, ok := anything.(string)
	if ok {
		fmt.Println("anythingは文字列型です:", str)
	} else {
		fmt.Println("anythingは文字列ではありません")
	}

	// 型アサーションの例2: インターフェース型の値を整数型にキャストする
	num, ok := anything.(int)
	if ok {
		fmt.Println("anythingは整数型です:", num)
	} else {
		fmt.Println("anythingは整数型ではありません")
	}

	// 型アサーションの例3: インターフェース型の値をスライス型にキャストする
	slice, ok := anything.([]int)
	if ok {
		fmt.Println("anythingはスライスです:", slice)
	} else {
		fmt.Println("anythingはスライスではありません")
	}
}
