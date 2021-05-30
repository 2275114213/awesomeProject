package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	// 排序
	nums := []int{3, 2, 1, 6, 9}
	sort.Ints(nums)
	fmt.Println(nums)
	names := []string{"a", "b", "d", "f", "e"}
	sort.Strings(names)
	fmt.Println(names)
	// 查找
	rand.Seed(time.Now().Unix())
	//num := rand.Intn(100)
	// [1,3,5,9,10]
	// x 是不是在切片中
	// 返回的是有序数组里面的可以插入的索引值
	res := sort.SearchInts(nums,5)
	fmt.Println(res)


}
