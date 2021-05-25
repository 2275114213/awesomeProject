package main

import "fmt"

func main() {
	var age int = 31 // 字面量  十进制 八进制 十六进制
	fmt.Printf("%T,%d,%#v\n", age, age, age)
	/*
		8进制：0？？？ <8
		16进制：0X？？<16 0-9A-F
		2进制 ：
	*/
	//	常用操作
	//  算数运算：+ - * / % ++ --
	a, b := 2, 4
	fmt.Println(a / b) // 0
	a++
	b--
	fmt.Println(a, b)
	c := 10
	d := 6
	fmt.Println(c / d) // 1 不会四舍五入
	// 关系运算 > < >= <=  == !=
	fmt.Println("f", a < b)
	fmt.Println("f", a > b)
	fmt.Println("t", a >= b)
	fmt.Println("t", a <= b)
	fmt.Println("t", a == b)
	fmt.Println("f", a != b)

	// 了解
	// 位运算 0111  （8421大法）
	// 按位与 &
	/*
		7   0111
		2   0010
			0010
		2
	*/
	// 按位或 ｜
	/*
				7 0111
				2 0010
		          0111
				7
	*/
	// 取非  ^
	fmt.Println("-3", ^2)

	/*
		7  0111  0010
		   1000  1101
		8
	*/
	// 右移位 >>
	/*
		7  0111>>2  = 0001
		   1000
		8
	*/
	// 左移位 <<
	// and not: &^  0110 &^ 1011 = 0100
	//1011 &^ 1101 = 0010
	//
	//&^ 二元运算符的操作结果是“bit clear"
	//a &^ b 的意思就是 清零a中，ab都为1的位
	// 0111  0010  0101=5
	var (
		achar   byte = 'a' //97
		Achar   byte = 'A' //65
		unicode rune = '中' //码点
	)
	fmt.Println(achar, Achar)
	fmt.Println(unicode)
}
