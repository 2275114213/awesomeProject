package main

import "fmt"

//import ("fmt","")  引入多个包

// 包级别
var packageVar string = "packageVar"

func main() {

	// 定义变量
	var msg string = "hello world"
	var name string = "hello world"
	shortSting := "short" //短声明一定要在函数级别，不能是包级别的
	fmt.Println(shortSting)
	var zeroString string   //定义类型但不初始化值
	var typeString = "zero" // 不定义类型但是需要初始化值
	fmt.Println(typeString)
	fmt.Println(zeroString)
	// 定义变量
	var hobby, mam = "sdds", "dsdff"
	fmt.Println("hobby", hobby, "mam", mam)
	// 函数级别
	var funcVar string = "funcVar"
	fmt.Println(funcVar, name)
	fmt.Println(name)
	fmt.Println(msg)
	{ //限制变量定义的
		// 块级别
		var blockVar string = "blockVar"
		fmt.Println("===", packageVar, funcVar, blockVar)
		{
			// 子块级别
			var innerBlockVar string = "inner bloack 4.var"
			fmt.Println(packageVar, funcVar, blockVar, innerBlockVar)
		}
	}
	// 其他代码
	fmt.Println(msg)
	// 其他代码
	var age int
	fmt.Println(age)
	var haha string
	fmt.Printf("%#v\n", haha)
	var test1 [2]string
	fmt.Println(test1)

	var test2 []string
	fmt.Println(test2 == nil)
	var test3 map[string]string
	fmt.Println(test3 == nil)
	test2 = make([]string, 2)
	test2[0] = "dsfdfd"

	myMap := make(map[string]string) //myMap2 map[string]string{}
	fmt.Println(myMap)
	var myMap2 map[string]string
	fmt.Println(myMap2 == nil)

	var myMap3 map[string]string = map[string]string{"1": "L"}
	fmt.Println(myMap3)
	for key, value := range myMap3 {
		println(key, value)
	}
}
