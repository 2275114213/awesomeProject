package main

import "fmt"

func main() {
	name := ""
	fmt.Print("请输入你的名字")
	fmt.Scan(&name)
	fmt.Println("请输入你的名字", name)
	age := 0
	fmt.Print("请输入你的年龄")
	fmt.Scan(&age)
	fmt.Println("请输入你的年龄", age)
	// 如果输入其他格式上面解析第一个字符报错，则不会要求输入msg，并把后面的字符赋值给 msg
	msg := ""
	fmt.Print("请输入你的msg")
	fmt.Scan(&msg)
	fmt.Println("请输入你的msg", msg)

}
