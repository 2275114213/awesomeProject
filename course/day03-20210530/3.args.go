package main

import "fmt"

func paramTest(args ...string) {
	fmt.Println(args) // args 为切片类型
}
func paramTest1(args ...string) {
	paramTest(args...) // args 为切片类型,把切片值传递给函数
}

// 1. 可变参数在一个方法中只能有一个，并且放在函数变量声明最后
func main() {
	paramTest("sdsdsfds")
	paramTest1("sdsdsfds")
}
