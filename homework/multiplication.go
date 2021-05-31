package main

import "fmt"

/*
1*1=1
1*2=2 2*2=4
1*3=3 2*3=6 3*3=9

*/

func main() {
	// fmt.Printf("[%010d]", 1000)
	// 占10个，默认不够空格填充，现在是按0填充
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%2d * %2d = %2d    ", i, j, i*j)
		}
		fmt.Println()
	}
}
