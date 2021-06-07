package main

import "fmt"

// interface 本质是一个指针
type Animal interface {
	Sleep()
	GetColor() string // 获得动物颜色
	GetType() string
}

// 具体的类
type Cat struct {
	color string
}

func (t *Cat) Sleep() {
	fmt.Println("Cat is sleep")
}
func (t *Cat) GetColor() string {
	fmt.Println("Cat is sleep")
	return t.color
}

func (t *Cat) GetType() string {
	fmt.Println("Cat is sleep")
	return "Cat"
}

type Dog struct {
	color string
}

func (t *Dog) Sleep() {
	fmt.Println("Dog is sleep")
}
func (t *Dog) GetColor() string {
	fmt.Println("Dog is sleep")
	return t.color
}

func (t *Dog) GetType() string {
	fmt.Println("Dog is sleep")
	return "Cat"
}

func showAnimal(animal Animal) {
	animal.Sleep()
}
func main() {
	//var animal Animal // 接口的数据类型，父类指针
	//animal = &Cat{"Green"}
	//animal.Sleep() // 调用的就是Cat 的Sleep()
	//animal = &Dog{"White"}
	//animal.Sleep()
	showAnimal(&Cat{"Green"})
}
