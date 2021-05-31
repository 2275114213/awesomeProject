package main

import "fmt"
func cal1(a, b int, c func(d int, f int) int) (r2 int, ) {
	r2 = c(a, b)
	return r2
}
func main() {
	// 匿名函数
	add1 := func(n1 int, n2 int) int {
		return n1 + n2
	}
	mutil := func(n1 int, n2 int) int {
		return n1 * n2
	}
	rt := add1(1, 2)
	rt1 := mutil(1, 2)
	fmt.Println(rt)
	fmt.Println(rt1)

	res := cal1(10, 20, func(n1, n2 int) int {
		return n1 * n2
	})
	fmt.Println(res)

}
