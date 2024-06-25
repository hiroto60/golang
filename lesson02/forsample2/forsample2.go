package main

import "fmt"

func main() {
	numbers := [5]int{50, 12, 39, 77, 65}

	for index, value := range numbers {
		fmt.Printf("index: %d, value: %d\n", index, value)
	}

	for _, value := range numbers {
		fmt.Printf("value: %d\n", value)
	}
}
