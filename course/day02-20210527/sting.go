package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

func main() {
	ascii := "abc"
	fmt.Println([]byte(ascii))
	asciiChina := "abc我爱中国"
	// 转成bytes
	fmt.Println([]byte(asciiChina))
	fmt.Println([]rune(asciiChina))
	fmt.Println(utf8.RuneCountInString(asciiChina))
	fmt.Println(string([]byte(asciiChina)))
	fmt.Println(string([]rune(asciiChina)))
	// 字符串转化成int。将字符串转换为 int 类型
	// Atoi 相当于 ParseInt(s, 10, 0)
	fmt.Println(strconv.Atoi("1999"))
	// int 转化成字符串
	fmt.Println(strconv.Itoa(12))
	// 字符串转换为布尔值
	// 它接受真值：1, t, T, TRUE, true, True
	// 它接受假值：0, f, F, FALSE, false, False.
	// 其它任何值都返回一个错误
	// FormatBool 将布尔值转换为字符串 "true" 或 "false"
	fmt.Println(strconv.ParseBool("True"))
	// ParseFloat 将字符串转换为浮点数
	// s：要转换的字符串
	// bitSize：指定浮点类型（32:float32、64:float64）
	// 如果 s 是合法的格式，而且接近一个浮点值，
	// 则返回浮点数的四舍五入值（依据 IEEE754 的四舍五入标准）
	// 如果 s 不是合法的格式，则返回“语法错误”
	// 如果转换结果超出 bitSize 范围，则返回“超出范围”

	fmt.Println(strconv.FormatFloat(12.1213,'f',6,64))
	fmt.Println(strconv.ParseFloat("12.1213",64))

}
