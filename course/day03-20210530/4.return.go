package main

import (
	"fmt"
)

func addTest(n1, n2 int) (int, int) {
	return n1 + n2, n1 * n2
}

// 命名返回值
func add1(n1, n2 int) (r1 int, r2 int) {
	r1 = n1 + n2
	r2 = n1 * n2
	return
}

func add3(n1, n2 int) (r1, r2 int) {
	r1 = n1 + n2
	r2 = n1 * n2
	return
}
func main() {
	a, b := addTest(1, 3)
	fmt.Println(a, b)
	a, b = add1(1, 3)
	fmt.Println(a, b)
	a, b = add3(1, 3)
	fmt.Println(a, b)
}
