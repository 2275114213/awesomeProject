package main

import "fmt"

func main() {
BREAK:
	for i := 0; i < 3; i++ {
		for j := 0; j <= 5; j++ {
			fmt.Println(j)
			if j == 2 {
				break BREAK
			}
		}
	}

}
