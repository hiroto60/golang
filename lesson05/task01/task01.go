package main

import "fmt"

type Pet interface {
	Name() string
	Eat()
	Play()
	Sleep()
}

type Dog struct {
	name   string
	energy int
}

type Cat struct {
	name string
	mood string
}

// dog
func (d *Dog) Name() string {
	return d.name
}

func (d *Dog) Eat() {
	d.energy += 10
	fmt.Printf("%vは食事をしています。エネルギー:%v\n", d.name, d.energy)
}

func (d *Dog) Play() {
	d.energy -= 5
	fmt.Printf("%vは遊んでいます。エネルギー:%v\n", d.name, d.energy)
}

func (d *Dog) Sleep() {
	d.energy += 20
	fmt.Printf("%vは眠っています。エネルギー:%v\n", d.name, d.energy)
}

// cat
func (c *Cat) Name() string {
	return c.name
}

func (c *Cat) Eat() {
	c.mood = "happy"
	fmt.Printf("%vは食事をしています。気分:%v\n", c.name, c.translateMood())
}

func (c *Cat) Play() {
	c.mood = "playful"
	fmt.Printf("%vは遊んでいます。気分:%v\n", c.name, c.translateMood())
}

func (c *Cat) Sleep() {
	c.mood = "sleepy"
	fmt.Printf("%vは眠っています。気分:%v\n", c.name, c.translateMood())
}

func (c Cat) translateMood() string {
	switch c.mood {
	case "happy":
		return "幸せ"
	case "playful":
		return "遊び心"
	case "sleepy":
		return "眠い"
	}

	return "不明"
}

func main() {
	// ペットのスライスを作成し、犬と猫を追加します
	pets := []Pet{
		&Dog{name: "バディ", energy: 50},
		&Cat{name: "ウィスカー", mood: "happy"},
	}

	// ペットのスライスを反復処理して、それぞれのペットが食事をして、遊んで、眠る様子を表示します
	for _, pet := range pets {
		fmt.Printf("%vの情報:\n", pet.Name())
		pet.Name()
		pet.Eat()
		pet.Play()
		pet.Sleep()
	}
}
