package main

import "fmt"

func main() {
	// nilMap 报错
	var nilMap map[string]string = map[string]string{}
	socre := map[string]string{}
	nilMap["kk"] = "teacher"
	fmt.Println(nilMap)
	fmt.Println(socre)
}
