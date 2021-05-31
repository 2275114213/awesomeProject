package main

import "fmt"

func addBase(base int) func(int) int {
	fmt.Println(base)
	return func(n int) int {
		return n + base
	}
}
func main() {
	add1 := addBase(1)
	fmt.Println(add1(10))
}
