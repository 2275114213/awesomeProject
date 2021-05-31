package main

import "fmt"

// 值类型,在函数内修改实参的值
// fmt.Scan(&age)

// 指针传递，传递的指针，*point 解引用
func changePointer(point *int) {
	*point++
	//point++ 不能对内存地址进行操作，对内存中的值进行操作
}
func change(point int) {
	point++
}
func main() {
	age := 65
	change(age)
	fmt.Println(age)
	changePointer(&age)
	fmt.Println(age)


}
