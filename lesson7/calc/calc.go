package calc

import "fmt"

// 2つの整数の加算を行う
func Add(a, b int) int {
	return a + b
}

// 2つの整数の減算を行う
func Subtract(a, b int) int {
	return a - b
}

// 2つの整数の乗算を行う
func Multiply(a, b int) int {
	return a * b
}

// 2つの整数の除算を行う
// ゼロ除算が発生した場合、エラーを返す
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}
