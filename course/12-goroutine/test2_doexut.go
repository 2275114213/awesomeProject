package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		// 用go 创建承载一个形参为空, 返回值为空的一个函数
		go func() {
			defer fmt.Println("A.defer")
			func() {
				defer fmt.Println("B.defer")
				fmt.Println("B")
				runtime.Goexit()
			}()
			fmt.Println("A")
		}()
		// 死循环
		for {
			time.Sleep(1 * time.Second)
		}
	*/
	go func(a, b int) bool {
		fmt.Println("a=", a, "b=", b)
		return true
	}(10, 20)
	for {
		time.Sleep(1 * time.Second)
	}
}
