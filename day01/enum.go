package main

import "fmt"

func main() {
	const (
		Mon = iota // 0 ,在常量组里面使用，每次调用加+1
		B        // 1
		C        // 2
		D = iota
	)
	fmt.Println(Mon, B, C,D)
}
