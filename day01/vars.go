package main

import "fmt"

//import ("fmt","")  引入多个包

// 包级别
var packageVar string = "packageVar"

func main() {

	// 定义变量
	var msg string = "hello world"
	var name string = "hello world"
	// 函数级别
	var funcVar string = "funcVar"
	fmt.Println(funcVar, name)
	fmt.Println(name)
	fmt.Println(msg)
	{
		// 块级别
		var blockVar string = "blockVar"
		fmt.Println("===", packageVar, funcVar, blockVar)
		{
			var innerBlockVar string = "inner bloack var"
			fmt.Println(packageVar, funcVar, blockVar, innerBlockVar)
		}
	}
	// 其他代码
	fmt.Println(msg)
	// 其他代码
}
