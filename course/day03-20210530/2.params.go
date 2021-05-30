package main

import "fmt"

// 连续多个变量的类型相同
//func add(n1 int, n2 int) {
//	fmt.Println(n1, n2)
//}

func add(n1, n2 int) {
	fmt.Println(n1, n2)
}
func add2(n1, n2 int, s1, s2, s3 string) {
	fmt.Println(n1, n2, s2, s3, s1)
}

func main() {
	add(1, 2)
	add2(1, 2, "sds", "sds", "sdsd")
}
