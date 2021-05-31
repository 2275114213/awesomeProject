package main

import (
	"fmt"
)

func main() {
	y := ""
	fmt.Print("看到卖西瓜的")
	fmt.Scan(&y)
	switch y {
	case "yes", "y", "Y":
		fmt.Println("买1个")
	case "no", "n", "N":
		fmt.Println("买10个")
	default:
		fmt.Println("买10个")

	}
	//if y == "yes" {
	//	fmt.Println("买一个包子")
	//} else {
	//	fmt.Println("买10个包子")
	//}
}
