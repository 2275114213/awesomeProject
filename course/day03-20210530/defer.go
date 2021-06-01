package main

import "fmt"

/*
defer是在return之前执行的。这个在 官方文档中是明确说明了的。要使用defer时不踩坑，
最重要的一点就是要明白，return xxx这一条语句并不是一条原子指令!
函数返回的过程是这样的：先给返回值赋值，然后调用defer表达式，最后才是返回到调用函数中。
*/
//func f() (r int) {
//	defer func(r int) {
//		r = r + 5
//	}(r)
//	return 1
//}
func f() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}
func testDefer() (rt string) {
	// 延迟函数执行, 在函数退出之前执行,多个defer 堆栈执行
	defer func() () {
		fmt.Println("test API")
		rt = "defer"
	}()

	defer func() {
		fmt.Println("test API1")
	}()
	defer func() {
		fmt.Println("test API2")
	}()
	fmt.Println("main")
	return rt
}
func main() {
	// 延迟函数执行, 在函数退出之前执行,多个defer 堆栈执行
	res := testDefer()
	fmt.Println(res)
	fmt.Println(f())
}
