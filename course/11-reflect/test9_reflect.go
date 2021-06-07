package main

import (
	"fmt"
	"reflect"
)

func reflectNum(i interface{}) {
	fmt.Println("type", reflect.TypeOf(i))
	fmt.Println("type", reflect.ValueOf(i))
}

func main() {
	var num float64 = 1.2345
	reflectNum(num)
}
