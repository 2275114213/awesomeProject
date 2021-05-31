package main

import (
	"fmt"
	"strings"
)



func taskPrint(task map[string]string) {
	fmt.Println(task)
}

func main() {
	var text string
	fmt.Println("请输入任务名称")
	fmt.Scan(&text)
	for _, todo := range todos {
		if strings.Contains(todo["name"], text) {
			fmt.Println(strings.Repeat("-", 20))
			taskPrint(todo)
		}
	}
}
