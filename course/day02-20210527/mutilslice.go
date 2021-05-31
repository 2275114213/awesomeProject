package main

import "fmt"

func main() {
	var names = [][]string{}
	names = [][]string{[]string{"a", "b"}, []string{"2", "3"}, []string{"da", "sd"}}
	names = append(names, []string{"3", "5"})
	fmt.Println(names)
	fmt.Println(names[0][1])
}
