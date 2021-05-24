package main

import (
	"fmt"
	"strconv"
)

func main() {
	// int int8 type()
	// float32 float64 type()
	// string
	var (
		intVal     = 1
		float64Val = 2.2
		stringVal  = "3"
	)
	fmt.Println(intVal, float64Val, stringVal)
	fmt.Printf("%T,%#v\n", float64(intVal), float64(intVal))
	fmt.Printf("%T,%#v\n", string(intVal), string(intVal))
	// string 不能直接转化int,借助语言包
	v, err := strconv.Atoi(stringVal)
	fmt.Println(v, err)
	//定义变量变量名不能相同
	// 转化为浮点数
	vv, err := strconv.ParseFloat(stringVal, 64)
	fmt.Println(vv, err)
	// 转化为整数
	vvv, err := strconv.Atoi(stringVal)
	fmt.Println(vvv, err)


}
