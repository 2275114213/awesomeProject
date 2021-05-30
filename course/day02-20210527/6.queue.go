package main

import (
	"fmt"
)

func main() {
	// 队列 先进先出
	queue := []string{}
	queue = append(queue, "x", "y", "z")
	fmt.Println(queue)
	x := queue[0]
	queue = queue[1:]
	fmt.Println(x)
	y := queue[0]
	queue = queue[1:]
	fmt.Println(y)
	z := queue[0]
	queue = queue[1:]
	fmt.Println(z)

}
