package main // 当前程序的包

import (
	"fmt"
	t "time" // 给包起个别名
)

//程序入口
func main() {
	// goland 中的表达式加；和不加都可以
	fmt.Println("hello world")
	t.Sleep(1 * t.Second)
}

// go fmt 单行注释
/*
块注释
我是一个注释
*/
