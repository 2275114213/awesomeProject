package main

import "fmt"

func main() {
	name, desc := "kk", "sds"
	func(name string) {
		name = "sdsd"
		//修改的是上一个级别的
		desc = "sdsd"
		fmt.Println(name, desc)
	}("sdsd")
	fmt.Println(name, desc)
}
