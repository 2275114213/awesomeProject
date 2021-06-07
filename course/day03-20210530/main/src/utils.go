package main

import "fmt"

var utilsVar = "utils Var"

func utilsFunc() {
	fmt.Println("utils Func")
}

// GO PATH 项目
// GOPATH GOMODULE
// 1。同一文件夹下所有go 文件的包名必须一致
// 2。关闭GOMODULE
// 3。在项目目录直接运行go build 编译，将当前文件夹下的文件进行编译
// main 包中只能有一个main函数

/*
GOPATH 环境变量信息，定义多个目录
	src == 》 源文件
	pkg == 》 包文件
	bin == 》 程序编译的可执行文件

编译程序 go build 项目文件路径(src)


*/
