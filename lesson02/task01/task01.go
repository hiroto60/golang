package main

import "fmt"

func main() {

	var name string = "山﨑光都"
	var age int = 25
	var birthPlace string = "東京都"
	var hobby string = "超ときめき宣伝部(推しのアイドル)"

	fmt.Printf("はじめまして、%sと申します。\n年齢は%v歳で、%s出身です.\n趣味は%sです。\nよろしくお願いします。",
		name, age, birthPlace, hobby)
}
