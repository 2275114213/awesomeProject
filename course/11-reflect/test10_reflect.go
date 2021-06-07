package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (t User) Call() {
	fmt.Println("user is call")
	fmt.Printf("%#v\n", t)
}

func main() {
	user := User{Id: 1, Name: "lxl", Age: 1}
	user.Call()
	DoFieldAndMethod(user)

}

func DoFieldAndMethod(input interface{}) {
	// 获取input 的 value
	inputValue := reflect.ValueOf(input)
	fmt.Println(inputValue)
	// 获取input 的 type
	inputType := reflect.TypeOf(input)
	//fmt.Println(inputType)
	//fmt.Println(inputType.Name())
	// 分别通过type 获取里面的字段
	// 1. 获取interface 的 11-reflect.TypeOf， 通过Type 得到NumField
	// 2. 得到每个field ，数据类型
	// 3. 通过field 又一个InterFace() 方法得到对应的value
	// 	t := 11-reflect.TypeOf(point).Elem()  传地址用这个
	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		value := inputValue.Field(i)
		sdsds := inputValue.Field(i).Interface()
		fmt.Println("======", field.Name, field.Type, value, sdsds)
	}
	// 通过Type 获取里面的方法
	for j := 0; j < inputType.NumMethod(); j++ {
		m := inputType.Method(j)
		fmt.Println("======", m)
	}
}
