package main

import "fmt"

func main() {
	// 函数内定义的变量一定要使用
	var name = "ok"
	var age int
	hobby, comment := "sleep", "sdsdw"
	var zreo, country = "bane", "asd"
	fmt.Println(name, age, hobby, zreo, country, comment)
	var (
		name1 string = "name"
		msg   string = "mag"
		desc  string
	)
	{
		name = "sdsd"
		fmt.Println(name)
	}
	fmt.Println(name, name1, msg, desc)
}
