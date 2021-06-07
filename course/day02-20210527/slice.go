package main

import "fmt"

func name1(name map[string]string) {
	name["ads"] = "sdsd"
}

// 切片：长度可变 相同类型的一组值的序列
// 类型：[]type
func main() {
	var names []string
	fmt.Printf("%T,%#v\n", names, names) // nil
	// 初始化
	// 字面量
	names = []string{} // 空切片，已经初始化，但元素为0
	fmt.Printf("%T,%#v\n", names, names)
	names = []string{"name1", "name2"}
	fmt.Printf("%T,%#v\n", names, names)

	//  索引
	names = []string{1: "sd", 2: "sd", 100: "sds"}
	fmt.Printf("%T,%#v,%v\n", names, names, len(names))

	// make 参数

	names = make([]string, 5) // 2个参数,申请有五个字符串元素的切片
	fmt.Println(names)
	fmt.Println(cap(names))
	fmt.Println(len(names))
	//make() // 3个参数
	names = make([]string, 0, 10) // 空切片，但是容量是10
	fmt.Printf("%#v\n", names)
	fmt.Println(len(names))
	// 容量

	// 添加元素
	fmt.Println(cap(names))
	names = append(names, "sds")
	fmt.Println(names)
	fmt.Printf("%#v\n", names)
	fmt.Println(len(names))
	fmt.Println(cap(names))
	// 遍历数组
	for i, v := range names {
		fmt.Println(i, v)
	}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	// copy 切片之间赋值
	aSLise := []string{"a", "b", "c", "e", "f"}
	bSlice := []string{"d", "e", "f", "g"}
	// 目标，源, 源多会删除源元素,目标多会保留目标元素

	copy(aSLise, bSlice)

	fmt.Println(aSLise)
	fmt.Println(bSlice)

	// 切片操操作
	// 字符串切片： msg[start:end]

	nums := []int{1, 2, 34} //len=3 cap=3
	numChild := nums[1:2]   // start <= end<=cap   new_cap = cap-start
	fmt.Println(numChild)
	test_name := make(map[string]string)
	test_name["name"] = "sdsd"
	name1(test_name)
	fmt.Println(test_name)
}
