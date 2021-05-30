// 包的声明
package main

import "fmt"

// 第三方标准包，自定义
// 包级别的变量，常量， 函数

// 无参数 无返回值
func sayHello() {
	fmt.Println("Hello World")
}

// 有参数
func sayHello1(name string, name1 string) {
	fmt.Println("Hello World", name, name1)
}

// 有参数
func sayHello2(n int, n1 int) int {
	return n + n1
}

func test(n int, n1 string) {
	fmt.Println("Hello World", n, n1)

}
func main() {
	// 调用 方法名()
	sayHello()
	sayHello1("刘晓蕾", "刘晓寒")
	res := sayHello2(1, 2)
	fmt.Println(res)
	test(1, "sdd")
}
