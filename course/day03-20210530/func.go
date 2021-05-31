package main

import "fmt"

// 无参数 无返回值
func sayHello13() {
	fmt.Println("Hello World")
}

// 有参数
func sayHello14(name string, name1 string) {
	fmt.Println("Hello World", name, name1)
}

func sayHello15(name int, name1 int) int {
	return name + name1
}

func add10(n1, n2 int) int {
	return n1 + n2

}
func mutil1(n1, n2 int) int {
	return n1 * n2

}
func cal(a, b int, c func(d int, f int) int) (r2 int, ) {
	r2 = c(a, b)
	return r2
}

func main() {
	// 定义变量为函数类型
	var callbackFunc func(int, int) int
	callbackFunc = sayHello15
	res := callbackFunc(1, 2)
	sayHello()
	sayHello1("sds","sdsd")
	fmt.Println("============", cal(10, 20, add10))
	fmt.Println(res)
	fmt.Printf("%T,#%v\n", callbackFunc, callbackFunc)
	fmt.Printf("%T,#%v\n", callbackFunc, callbackFunc)
	fmt.Printf("%T\n", sayHello)
	fmt.Printf("%T\n", sayHello1)
	sayHello1("liuxiailei", "liuxiaohan")

	// 匿名函数
	func(){
		fmt.Println("test")
	}()

}
