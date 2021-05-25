package main

import "fmt"

func main() {
	y := ""
	fmt.Print("看到卖西瓜的")
	fmt.Scan(&y)
	if y == "yes" {
		fmt.Println("买一个包子")
	} else {
		fmt.Println("买10个包子")
	}
}
