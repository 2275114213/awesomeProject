package main

import (
	"fmt"
	"reflect"
)

type resume struct {
	Name string `info:"name" doc:"我的名字"`
	Sex  int    `info:"sex"`
}

func findTag(point interface{}) {
	t := reflect.TypeOf(point).Elem()
	for i := 0; i < t.NumField(); i++ {
		info := t.Field(i).Tag.Get("info")
		fmt.Println(info)
	}
}
func main() {
	var re resume
	re = resume{Name: "sdsd", Sex: 1}
	findTag(&re)
}
