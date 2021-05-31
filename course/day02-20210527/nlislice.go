package main

import "fmt"

func main() {
	var nilSlice []int  // nli 切片
	emptySlice := []int{} // 空切片
	fmt.Printf("%#v\n", emptySlice)
	fmt.Printf("%#v\n", nilSlice)
}
