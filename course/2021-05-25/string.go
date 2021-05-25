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
	fmt.Printf("%T,%#v,%c\n", msg[0], msg[0], msg[0])
	fmt.Println(msg[1:4])
	// len 是字节的大小
	fmt.Println(len(msg))
	fmt.Println(len(msgRaw))
}
