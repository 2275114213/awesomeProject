package main

import (
	"fmt"
)

func main() {
	//score := 0
	var score float32
	fmt.Println("请输入你的分数")
	fmt.Scan(&score)
	if score >= 90 {
		fmt.Println("A")
	} else if score >= 80 {
		fmt.Println("B")
	} else if score > 60 {
		fmt.Println("C")
	} else {
		fmt.Println("D")
	}

}
