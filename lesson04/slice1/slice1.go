package main

import "fmt"

func main() {
	// makeでスライスを作成
	slice1 := make([]int, 5, 10)
	fmt.Println(slice1) // [0 0 0 0 0]

	// スライスリテラルで作成
	slice2 := []int{1, 2, 3, 4, 5}
	fmt.Println(slice2) // [1 2 3 4 5]

	// 配列からスライスを作成
	arr := [5]int{1, 2, 3, 4, 5}
	slice3 := arr[1:4]
	fmt.Println(slice3) // [2 3 4]

	// 既存のスライスから新しいスライスを作成
	subSlice1 := slice2[2:]
	fmt.Println(subSlice1) // [3 4 5]
	subSlice2 := slice2[:3]
	fmt.Println(subSlice2) // [1 2 3]
}
