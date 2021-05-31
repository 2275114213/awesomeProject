package main

import "fmt"

// 赋值 《==》 实参/形参
func test1(n int) {
	n = 1
}
func test2(s []int) {
	s[0] = 1
}

func main() {
	//age := 30
	//tmpAge := age
	//tmpAge = 31
	//// age 不会变（在内存中申请新的空间，将a的值拷贝带b中）
	//// 在修改锕不影响b
	//fmt.Println(tmpAge, age)
	//
	//users := []string{"asd", "sdsds"}
	//tmpUsers := make([]string, 10)
	//tmpUsers = users
	//tmpUsers[0] = "kk"
	//// users 会变
	//fmt.Println(users, tmpUsers)
	//pointA := &age
	//pointB := pointA
	//*pointB = 32
	//fmt.Println(*pointA, *pointB)
	a := 1
	b := make([]int, 10)
	test1(a)
	test2(b)
	fmt.Println(a)
	fmt.Println(b)
}
