package main

import "fmt"

/*
选班长： 不记名投票
*/
// 统计"我有一个梦想"中每个英文(大小写区分)字母出现的次数
func main() {
	var desc string = "I have a dream"
	// 过滤字符，整数
	// a-z A-Z

	var countMap = map[rune]float64{}
	for _, i := range desc {
		countMap[i] ++
	}
	fmt.Println(countMap)
}
