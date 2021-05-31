package main

import (
	"fmt"
	"sort"
)

func main() {
	stats := [][]int{{'A', 4}, {'B', 3}, {'C', 5}}
	// 使用出现次数进行排序
	sort.Slice(stats, func(i, j int) bool { return stats[i][1] > stats[j][1] })
	fmt.Println(stats)

	// 升序 <=
	// 降序 >=
	index := sort.Search(len(stats), func(i int) bool {
		return stats[i][1] == 3
	})
	fmt.Println(index)
}
