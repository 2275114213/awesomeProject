package main

import (
	"fmt"
	"strconv"
)



func genId() int {
	var rt int
	for _, todo := range todos {
		todoId, _ := strconv.Atoi(todo["id"])
		if rt < todoId {
			rt = todoId
		}
	}
	return rt + 1
}
func newTask(title, start_time, statusCode, people string) map[string]string {
	var task = make(map[string]string)
	task[name] = title
	task["id"] = strconv.Itoa(genId())
	task[startTime] = start_time
	task[user] = people
	task[status] = "未执行"
	return task
}
func main() {
	var title string
	var start_time string
	var status_code string
	var people string
	//var task = map[string]string{}
	// map 零值是不能操作的
	fmt.Println("请输入任务信息")
	fmt.Print("请输入任务名")
	fmt.Scan(&title)
	fmt.Print("请输入开始时间")
	fmt.Scan(&start_time)
	fmt.Print("请输入负责人")
	fmt.Scan(&people)
	task := newTask(title, start_time, status_code, people)
	todos = append(todos, task)
	fmt.Println(todos)

}
