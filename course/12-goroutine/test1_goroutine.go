package main

import (
	"fmt"
	"time"
)

func newTask() {
	i := 0
	for {
		i++
		fmt.Printf("New  Goroutine:i = %d\n", i)
		time.Sleep(1 + time.Second)
	}
}

// 主 12-goroutine
func main() {
	// 创建一个goroutine 执行newTask() 流程
	go newTask()
	//i := 0
	/*
	for {
		i++
		fmt.Printf("main 12-goroutine: i =%d\n", i)
		time.Sleep(1 * time.Second)
	}

	 */
}
