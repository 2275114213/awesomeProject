package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 3) //带有缓冲的channel

	go func() {
		defer fmt.Println("goroutine 结束")
		for i := 0; i < 4; i++ {
			c <- i
			fmt.Println("go runing", i, len(c), cap(c))
		}
	}()
	time.Sleep(2 * time.Second)
	for i := 0; i < 4; i++ {
		num := <-c
		fmt.Println(num)

	}
	fmt.Println("main 结束")
}
