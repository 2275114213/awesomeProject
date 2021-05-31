package main

import "fmt"

func main() {
	// map 映射 key-value 的数据结构
	// 无序的数据结构
	// 类型 ：map[KeyType]valueType
	var scores map[string]float64  //nil
	fmt.Printf("%T,%v\n", scores, scores)

	// 初始化
	// 字面量
	scores = map[string]float64{} // 空的map
	fmt.Printf("%T,%#v\n", scores, scores)

	scores = map[string]float64{"22": 89, "23": 90}
	fmt.Printf("%T,%#v\n", scores, scores)

	// make

	scores = make(map[string]float64) // == scores = map[string]float64{}
	scores = map[string]float64{"12": 90, "xx": 0}
	// 操作
	fmt.Println(len(scores))

	// 查
	fmt.Println(scores["12"])
	fmt.Println(scores["xx"]) // key 不存在，返回对应类型的零值

	// 判断key 是否存在
	v, ok := scores["yy"]
	fmt.Println(v, ok)

	v, ok = scores["xx"]
	fmt.Println(v, ok)
	// 增加
	scores["asa"] = 90
	fmt.Println(scores)
	// 删除
	scores["xx"] = 100
	fmt.Println(scores)
	//修改
	delete(scores, "xx")
	fmt.Println(scores)

	// 遍历,默认一个值为key
	for i, j := range scores {
		fmt.Println(i, j)
	}
}
