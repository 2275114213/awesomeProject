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
}
