package main

import "fmt"

func main() {
	var a string
	// pair<type:string,value:"aceld">
	a = "aceld"

	// pair<type:string,value:"aceld">
	var alkType interface{}
	alkType = a
	str, ok := alkType.(string)
	fmt.Println(alkType, str, ok)

}
