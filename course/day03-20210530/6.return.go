package main

import "fmt"

func testRecursion(flag bool) int {
	if flag {
		return 1
	} else {
		return 2
	}
}
func main() {
	res := testRecursion(true)
	fmt.Println(res)
}
