package main

import (
	"fmt"

)

// 做一个命令行的任务管理器
// 函数 ，输入输出，复合数据结构，基本数据类型
// 了解流程，对数据的操作流程，增删改查
// 1. 添加任务
// 2. 任务列表
// 3. 任务修改
// 4. 任务删除
// 5. 详情

// 任务任务名称，开始时间，结束时间，状态，负责人
// id,name,start_time,end_time,status,user
// []map[key]string
//var todos = make([]map[string]string, 0)
var todos = []map[string]string{
	{"id": "1", "title": "吃饭", "startTime": "startTime"},
}

const (
	name      = "title"
	startTime = "name"
	endTime   = "name"
	status    = "status"
	user      = "user"
)

const (
	statusNew = "未知性"
	complex   = "已完成"
)

func input(promt string) string {
	fmt.Print(promt)
	fmt.Scan(&promt)
	return promt
}

func main() {
	method := map[string]func(){
		"add":   todolist.add,
		"query": query,
	}
	for {
		text := input("请输入你的查询信息")
		if text == "add" {
			method, ok = method["add"]
			if method, ok = method["add"]; ok {
				method()
			}
		} else if text == "query" {
		} else {
		}
	}
}
