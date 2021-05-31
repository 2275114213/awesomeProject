package main

import "fmt"

// 数组：长度固定的存储的相同数据类型的一组值的序列
// 类型：[length]type
func main() {
	// 类型
	var names [3]string
	var signIns [3]bool
	var scores [3]float32
	//fmt.Println(names) 零值
	fmt.Printf("%T,%#v\n", names, names)
	fmt.Printf("%T,%#v\n", signIns, signIns)
	fmt.Printf("%T,%#v\n", scores, scores)
	// 字面量
	// 第一种
	names = [3]string{"new", "df", "sdds"}
	fmt.Printf("%#v", names)
	//names = [1]string{"new"}
	fmt.Printf("%#v\n", names)
	// 第二种，根据后面数据自动扩充
	age := [...]string{"new", "df"}
	fmt.Printf("%T\n", age)
	// 第三种
	second := [...]string{9: "index"}
	fmt.Printf("%#v\n", second)
	names = [3]string{1: "kk"}
	fmt.Printf("%#v\n", names)

	// 操作
	// 关系运算 == =
	fmt.Println(names == [3]string{})

	// 元素
	// 访问 &修改
	fmt.Printf("%#v\n", names[0])
	names[0] = "修改"
	fmt.Println(names)
	// 长度
	fmt.Println(len(names))
	// 遍历
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
	// range
	for i, v := range names {
		fmt.Println(i, v)
	}
	test := "sasd"
	for i := 0; i < len(test); i++ {
		// 打印的ascii码, println 打印就是字符形式
		fmt.Printf("%T,%#v,%c\n", test[i], test[i], test[i])
	}
	// 定义一个数组 每个元素也是一个数组
	// 二维数组
	d2 := [3][2]int{[2]int{1, 3}, [2]int{3, 4}, [2]int{1, 2}}
	fmt.Println(d2)
}
