package main

import (
	"fmt"
)

func main() {
	// z指针用来存放内存地址的变量
	var (
		pointerInt    *int
		pointerString *string
	)
	fmt.Println(pointerInt, pointerString)
	fmt.Printf("%T %#v\n", pointerInt, pointerInt)
	fmt.Printf("%T %#v\n", pointerString, pointerString)

	// 赋值
	// 使用现有变量取地址, &name
	age := 32
	age2 := age //age2 不会变
	fmt.Println("&age2", &age2)
	fmt.Printf("%T, %#v\n", &age, &age)
	pointerInt = &age
	fmt.Println(pointerInt)
	fmt.Println(*pointerInt)
	*pointerInt = 33
	age2 = 44
	fmt.Println(pointerInt, age, age2, &age2)

	// 另一种赋值
	pointerString = new(string)
	fmt.Printf("%T,%#v", pointerString, *pointerString)

}
