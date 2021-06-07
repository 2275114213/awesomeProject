package main

import "fmt"

type Human struct {
	name string
	age  int
}

func (t *Human) Eat() {
	fmt.Println("Human.Eat()...")
}
func (t *Human) Work() {
	fmt.Println("Human.Work()...")
}

type SuperMan struct {
	Human //superMan类继承Human属性
	level int
}

func (t *SuperMan) Eat() {
	fmt.Println("new superman")
}

func (t *SuperMan) Fly() {
	fmt.Println("new Work")
}
func main() {
	h := Human{name: "liuxiaolei", age: 1}
	h.Eat()
	h.Work()
	fmt.Println(h)
	// 方式一
	s := SuperMan{Human{"sdsd", 1}, 1}
	s.Eat()
	s.Work()
	s.Fly()
	// 方式二
	var sp SuperMan
	sp.name = " sdsd"
	sp.age = 1
	sp.level = 1
	sp.Fly()
}
