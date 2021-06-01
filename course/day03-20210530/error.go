package main

import (
	"errors"
	"fmt"
	"strconv"
)

func div(n1, n2 int) (int, error) {
	if n2 == 0 {
		return -1, errors.New("除数为0")
	} else {
		return n1 / n2, nil
	}
}

func main() {
	value, err := strconv.Atoi("xxxx")
	fmt.Println(value)
	fmt.Println(err)
	// 创建error
	e := fmt.Errorf("自定义错误")
	fmt.Printf("%T,%v", e, e) //*errors.errorString,自定义错误
	e2 := errors.New("自定义错误2")
	// 创建error
	fmt.Printf("%T,%v", e2, e2) //*errors.errorString,自定义错误
	// go语言希望内部如果有错误，通过最后一个返回值显示返回给调用者，由调用者决定如何处理

	if rt, err1 := div(1, 0); err1 == nil {
		fmt.Println("测试", rt)
	} else {
		fmt.Println("测试",err)
	}
}
