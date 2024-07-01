package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) SayHello() {
	fmt.Printf("こんにちは!私の名前は%sです。私は%d歳です。\n", p.Name, p.Age)
}

// 中の値を変更したい場合はポインタ型を使う
func (p *Person) IncrementAge() {
	p.Age++
	fmt.Println("誕生日を迎えました。")
}

func main() {
	p := Person{Name: "山田太郎", Age: 30}
	p.SayHello()     // 出力: こんにちは!私の名前は山田太郎です。私は30歳です。
	p.IncrementAge() // 出力: 誕生日を迎えました。
	p.SayHello()     // こんにちは!私の名前は山田太郎です。私は31歳です。
}
