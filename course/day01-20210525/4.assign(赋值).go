package main

import "fmt"

func main() {
	age := "12"
	fmt.Println(age)
	var name string = "haha"
	// 变量先定义在使用
	name = "dada"
	{
		name = "change name inner block" // 更改变量的值，不是定义变量
		fmt.Println(name)
		//4.var name = "change name inner block"  赋值
	}
	fmt.Println(name)
}
