package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		// 正确写法
		//func() {
		//	// 打开文件
		//	// 延迟关闭
		//	// 处理(处理出现错误)
		//}()
		fmt.Println("for before")
		// 错误写法，for循环之后才会执行defer
		defer func(i int) {
			fmt.Println("defer", i)
		}(i)
		fmt.Println("for after", i)
	}
	fmt.Println("main")
}
