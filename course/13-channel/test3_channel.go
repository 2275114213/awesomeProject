package main

import "fmt"

func main() {
	c := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		// close 可以关闭一个channel
		close(c)
	}()
	for {
		// 如果说ok 为true 表示channel 没有关闭或channel 有数据，如果为false 便是channel 已经关闭了
		if data, ok := <-c; ok {
			fmt.Println(data)
		} else {
			break
		}
	}
	fmt.Println("main finish")
}
