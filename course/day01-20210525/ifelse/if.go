package main

import "fmt"

func main() {

	fmt.Println("买10个包子")
	var y string
	fmt.Println("没有卖西瓜的")
	fmt.Scan(&y)
	fmt.Println("你输入的是", y)
	if y == "yes" {
		fmt.Println("买一个西瓜")
	} else {
		fmt.Println("不买西瓜")
	}

}
