package main

import "fmt"

func main() {
	subTotal := 1500

	tax := tax(subTotal)

	total := subTotal + tax

	fmt.Printf("1500円の商品の税込価格は%v円(消費税は%v円です。)", total, tax)
}

// 消費税額（商品価格の10%）を返す
func tax(value int) int {
	taxRate := 10

	return value / taxRate
}
