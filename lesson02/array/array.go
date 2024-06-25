package main

import "fmt"

func main() {

	var arr1 [5]int
	fmt.Println("arr1:", arr1)

	arr1[0] = 1
	arr1[1] = 2
	fmt.Println("arr1:", arr1)

	arr2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("arr2:", arr2)

	arr1Len := len(arr1)
	fmt.Println("arr1の長さ:", arr1Len)

}
