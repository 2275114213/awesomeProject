package main

import "fmt"

// interface 是万能数据类型
func myFunc(arg interface{}) {
	// interface 如何区分万能类型 此时的引用的底层数据类型到底是什么
	// 给 interface{} 提供'类型断言'的机制
	value, ok := arg.(map[string]string)
	fmt.Println("value", value, "ok", ok)
	fmt.Println("Myfunc is called")
	//fmt.Println(arg)
}

type Book1 struct {
	auth string
}

func main() {
	var book = Book1{
		auth: "sdsd",
	}
	myFunc(book)
	myFunc("book")
	myFunc(100)
	myFunc(100.12)
	people := make(map[string]string)
	people["name"] = "sdsd"
	people["sds"] = "sdsd"
	myFunc(people)
	fmt.Println(people)
}
