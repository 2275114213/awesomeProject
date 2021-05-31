package main

import "fmt"

func main() {
	total := 1
	index := 1
START:
	index += 1
	total += index
	if index == 100 {
		goto END //加到index==100结束
	}
	goto START
END:
	fmt.Println(total)
}
