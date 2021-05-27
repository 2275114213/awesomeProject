package main

import "fmt"

func main() {
	var (
		index = 0
		total = 0
	)
	// 如果没有break 就是死循环
	for {
		total += index
		index++
		if index > 100 {
			break
		}
	}
	fmt.Println(total, index)
}
