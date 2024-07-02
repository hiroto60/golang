package it

import (
	"fmt"
	"golang/lesson7/calc"
)

// 計算結果を保持
type Calculator struct {
	result int
}

// 新しいCalculatorを返す
func NewCalculator() *Calculator {
	return &Calculator{}
}

// 現在の計算結果にaを加算
func (c *Calculator) Add(a int) {
	c.result = calc.Add(c.result, a)
}

// 現在の計算結果からaを減算
func (c *Calculator) Subtract(a int) {
	c.result = calc.Subtract(c.result, a)
}

// 現在の計算結果にaを乗算
func (c *Calculator) Multiply(a int) {
	c.result = calc.Multiply(c.result, a)
}

// 現在の計算結果をaで除算
// 除算でエラーが発生した場合、エラーを返す
func (c *Calculator) Divide(a int) error {
	result, err := calc.Divide(c.result, a)
	if err != nil {
		return fmt.Errorf("division error: %w", err)
	}
	c.result = result
	return nil
}

// 現在の計算結果を返す
func (c *Calculator) Result() int {
	return c.result
}

// 計算結果を0にリセット
func (c *Calculator) Reset() {
	c.result = 0
}
