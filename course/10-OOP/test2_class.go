package main

import "fmt"

// 如果类名首字母大写，表示其他包也能够访问
type Hero struct {
	// 如果说类的属性首字母大写，表示该属性能对外开放，否则只能够类的内部访问
	Name  string
	Ad    int
	Level int
}

/*
func (t Hero) Show() {
	fmt.Println(t)
}

func (t Hero) GetName() string {
	return t.Name
}

func (t Hero) SetName(newName string) {
	// t 是调用该方法的对象的一个副本{拷贝}
	t.Name = newName
}
*/

func (t *Hero) Show() {
	fmt.Println(t)
}

func (t *Hero) GetName() string {
	return t.Name
}

func (t *Hero) SetName(newName string) {
	t.Name = newName
}
func main() {
	hero := Hero{Name: "zhangsan", Ad: 2, Level: 1}
	name := hero.GetName()
	hero.Show()
	hero.SetName("lisr")
	hero.Show()
	fmt.Println(name)
}
