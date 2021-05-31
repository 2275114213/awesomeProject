package main

import "fmt"

func main() {
	hights := []float32{12, 12.23, 544, 5454, 2324, 6767}
	for i := 0; i < len(hights)-1; i++ {
		for j := 0; j < len(hights)-1-i; j++ {
			if hights[j] > hights[j+1] {
				hights[j], hights[j+1] = hights[j+1], hights[j]
			}
		}
	}
	fmt.Println(hights)
}
