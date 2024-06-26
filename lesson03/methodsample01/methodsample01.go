package main

import "fmt"

func main() {
	// sumMethod1
	number1 := 100
	number2 := 10
	sumMethod1(number1, number2)

	// sumMethod2
	result2 := sumMethod2(number1, number2)
	fmt.Println("sumMethod2の結果は", result2)

}

func sumMethod1(num1 int, num2 int) {
	result := num1 + num2
	fmt.Println("sumMethod1の結果は", result)
}

func sumMethod2(num1 int, num2 int) int {
	result := num1 + num2
	return result
}
