package main

import "fmt"

func main() {
	// 控制台打印1。。。10
	// index 一般为赋值表达式，给控制变量赋初值；
	// condition： 关系表达式或逻辑表达式，循环控制条件；
	// condition： 关系表达式或逻辑表达式，循环控制条件；
	for index := 0; index < 10; index++ {

		fmt.Println(index)
	}
	// 计算1+2+3+4
	total := 0
	for index := 1; index <= 100; index++ {
		total += index
	}
	fmt.Println(total)

}
