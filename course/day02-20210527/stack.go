package main

import "fmt"

func main() {
	// 堆栈  先进后出
	stack := []string{}
	stack = append(stack, "x", "y", "z")
	fmt.Println(stack)
	x := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	fmt.Println(x)

	y := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	fmt.Println(y)

	z := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	fmt.Println(z)

}
