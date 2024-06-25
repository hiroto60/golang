package main

import "fmt"

func main() {
	for i := 2; i <= 100; i += 2 {
		if i > 20 {
			// 繰り返しから抜け出す
			break
		}
		fmt.Println(i)
	}
}
