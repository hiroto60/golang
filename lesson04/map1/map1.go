package main

import "fmt"

func main() {
	// makeでマップを作成
	map1 := make(map[string]int)
	map1["apple"] = 1
	map1["banana"] = 2
	fmt.Println(map1) // map[apple:1 banana:2]

	// makeでrune型のキーを持つマップを作成
	map2 := make(map[rune]int)
	map2['a'] = 1
	map2['b'] = 2
	fmt.Println(map2) // map[97:1 98:2]

	// マップリテラルで作成
	map3 := map[string]int{
		"apple":  1,
		"banana": 2,
		"orange": 3,
	}
	fmt.Println(map3) // map[apple:1 banana:2 orange:3]
}
