package main

import "fmt"

func main() {
	// 定一个channel
	c := make(chan int)
	go func() {
		defer fmt.Println("goroutine 结束")
		fmt.Println("goroutine 正在运行。。。。")
		c <- 666 // 将666发送给c，写到c 中
	}()

	num := <-c // 从管道c 中接受数据，并赋值给num
	fmt.Println(num)
	fmt.Println("main goroutine 结束")
}
