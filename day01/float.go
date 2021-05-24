package main

import "fmt"

func main() {
	var height float32 = 1.68
	fmt.Printf("%T %#v %f\n", height, height, height)
	var k float32 = 1e3 // 1*10^3
	fmt.Println(k)

	// 操作
	// 算数运算符 ： + - * / ++ --
	// 关系运算： > < >= <= 如果想要判断差值在一定区间范围内
	// 赋值运算 = =+ =/ =- =*
	height += k
	fmt.Println(height)
}
