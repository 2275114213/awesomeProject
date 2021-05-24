package main

import "fmt"

const packageName = "package"
const (
	packageMsg  = "msg"
	packageDesc = "desc"
)

func main() {
	const name string = "kk"
	const age = "12" // 常量可以不使用
	//name = "alias" 常量的值不可以被修改
	fmt.Println(name)
	const (
		A = "test"
		B // 使用前一个变量的值初始化
		C
		D = "D"
		E
		F
	)
	fmt.Println(A, B, C, D, E, D, F)
	fmt.Printf("%T, %v, %#v\n", A, A, A)
}
