package main

import "fmt"

func main() {
	age := 30
	tmpAge := age
	tmpAge = 31
	// age 不会变（在内存中申请新的空间，将a的值拷贝带b中）
	// 在修改锕不影响b
	fmt.Println(tmpAge, age)

	users := []string{"asd", "sdsds"}
	tmpUsers := make([]string, 10)
	tmpUsers = users
	tmpUsers[0] = "kk"
	// users 会变
	fmt.Println(users, tmpUsers)

}
