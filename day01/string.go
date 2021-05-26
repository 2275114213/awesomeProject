package main

import "fmt"

func main() {
	var msg = "你好\\n"
	var msgRaw = `你好\n` // 原生字符串
	fmt.Println(msg, msgRaw)
	// 操作
	// 字符串拼接 +
	fmt.Println(msg + msgRaw)
	fmt.Println(msg < msgRaw)
	fmt.Println(msg > msgRaw)
	fmt.Println(msg == msgRaw)
	msg += "我是刘晓蕾"
	fmt.Println(msg)
	// 切片 索引 ascii
	msg = "abcdefg"
	//整型   %c	该值对应的unicode码值
	//布尔型 %t	true或false
	//通用   %T	打印值的类型
	fmt.Printf("%T,%#v,%c\n", msg[0], msg[0], msg[0])
	fmt.Println(msg[1:4])
	// len 是字节的大小
	fmt.Println(len(msg))
	fmt.Println(len(msgRaw))

	// session 过期时间，及顺延方式
	// 应用与Cas集成
	// API 梳理
}
