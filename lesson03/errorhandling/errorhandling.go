package main

import (
	"fmt"
	"strconv"
)

func parseInt(str string) (int, error) {
	num, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}
	return num, nil
}

func main() {
	numStr := "123"
	num, err := parseInt(numStr)
	if err != nil {
		fmt.Println("数値の解析に失敗しました:", err)
		return
	}
	fmt.Println("数値:", num)

	numStr = "abc"
	num, err = parseInt(numStr)
	if err != nil {
		fmt.Println("数値の解析に失敗しました:", err)
		return
	}
	fmt.Println("数値:", num)
}
