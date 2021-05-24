package main

import "fmt"

func main() {
	//var isGirl bool
	isGirl := true
	fmt.Printf("%T,%#v\n", isGirl, isGirl)

	//与(&&)：左操作数与右操作数都为true ，结果为true
	a, b, c, d := true, true, false, false
	fmt.Println(a, b, c, d)
	fmt.Println("a&&b", a && b)
	fmt.Println("a&&c", a && c)
	fmt.Println("c&&d", c && d)
	//或(||)：左操作数与右操作数只要又一个为true，就为true
	fmt.Println("a||b", a || b)
	fmt.Println("a||c", a || c)
	fmt.Println("c||d", c || d)
	//非(!)：取反 true =》 false false=》true
	fmt.Println("!a", !a)
	fmt.Println("!c", !c)

	// 关系运算
	fmt.Println(a == b)
	fmt.Println(a == c)
	fmt.Println(a != c)
	fmt.Printf("%t", a)

}
