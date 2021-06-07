package main

import (
	"fmt"
)

// 声明一中行的数据类型myint 是int的一个别名
type myint int

func printBook(book Book) {
	// 传递book 的一个副本
	book.title = "测试"
	fmt.Println(book)
}


func changeBook(book *Book) {
	// 传递book 的一个副本
	book.title = "测试"
	fmt.Println(book)
}
// 定义一个结构体
type Book struct {
	title string
	auth  string
}

func main() {
	/*
		var a myint = 10
		fmt.Println(a)
	*/
	var book1 Book
	book1.title = "sf"
	book1.auth = "sf"
	printBook(book1)
	fmt.Printf("%#v\n", book1)
	changeBook(&book1)
	fmt.Printf("%#v\n", book1)

}
