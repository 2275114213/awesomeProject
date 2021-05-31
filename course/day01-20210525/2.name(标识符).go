package main

import "fmt"

func main() {
	/*

			标识符 ：程序中定义的名字，变量名，常量名字，函数名字，自定义类型，接口，包名字
			规范：
				1. 必须满足：组成只能非空的Unicode 编码字符串，数字，下划线组成
				2. 必须满足：变量名必须以unicode编码的字符串或下划线开头（不能以数字开头）
			    3. 必须满足：不能与go 关键字冲突
			建议：
				1. 建议：AsciII编码（a-z，A-Z）,数字，字母，下划线
				2. 变量使用驼峰式 myName
				3. 与go 内置标识符不要冲突（string ）
			说明：
				标识符区分大小写

		    内置常量：true，false，nil，iota（枚举类型）
			内置类型：bool，byte，rune，int，int8，int16，int32，int64，uint，uint8，uint16，uint32,uint64,uintptr,float32,float64,complex64,complex128,string,error
			内置函数：make，len,cap,new,append,copy,close,delete,complex,imag,panic,recover
			空白标识符： _
			关键字：
				声明：import package
				实体声明和定义：chan，const，func，interface，map，struct，type，4.var，
				流程控制：break，continue，case，default，defer，else，fallthrough，for，go,goto,if，range，return，select，switch
	*/
	var my = "my"
	var My = "My"
	fmt.Println(my, My)
}

//go build -o (其他文件) test.exe  day01-20210525/4.assign(赋值).go
