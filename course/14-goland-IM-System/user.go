package main

import (
	"fmt"
	"net"
	"strings"
)

type User struct {
	Name   string
	Addr   string
	C      chan string
	conn   net.Conn
	server *Server
}

// 创建一个用户的API
func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()
	user := &User{
		Name:   userAddr,
		Addr:   userAddr,
		C:      make(chan string),
		conn:   conn,
		server: server,
	}
	// 启动监听当前user channel 消息的 goroutine
	go user.ListenMessage()
	return user
}

// 监听当前User channel 的 方法,一旦有消息称就发送给客户端

func (t *User) ListenMessage() {
	for {
		message := <-t.C
		t.conn.Write([]byte(message + "\n"))

	}
}

// 上线业务
func (t *User) Online() {
	t.server.mapLock.Lock()
	// 注意是修改值，不是新的初始化
	t.server.OnlineMap[t.Name] = t
	t.server.mapLock.Unlock()
	// 广播当前用户上线的消息
	t.server.BroadCast(t, "已上线")
}

// 下线业务
func (t *User) Offline() {
	t.server.mapLock.Lock()
	// 注意是修改值，不是新的初始化
	delete(t.server.OnlineMap, t.Name)
	t.server.mapLock.Unlock()
	// 广播当前用户上线的消息
	t.server.BroadCast(t, "已下线")
}

func (t *User) SendMessage(msg string) {
	t.conn.Write([]byte(msg + "\n"))
}

// 发消息业务
func (t *User) DoMessage(msg string) {
	if msg == "who" {
		// 查询当前用户有哪些
		t.server.mapLock.Lock()
		for _, user := range t.server.OnlineMap {
			onlineMsg := fmt.Sprintf("[%s]%s:在线。。。\n", user.Addr, user.Name)
			t.SendMessage(onlineMsg)
		}
		t.server.mapLock.Unlock()
	} else if len(msg) > 7 && strings.HasPrefix(msg, "rename|") {
		// 当前消息格式为rename
		newName := strings.Split(msg, "|")[1]
		// 判断Name 是否已存在
		_, ok := t.server.OnlineMap[newName]
		if ok {
			t.SendMessage("已使用")
		} else {
			t.server.mapLock.Lock()
			delete(t.server.OnlineMap, t.Name)
			t.server.OnlineMap[newName] = t
			t.Name = newName
			t.server.mapLock.Unlock()
			t.SendMessage("更新成功")
		}

	} else if len(msg) > 4 && strings.HasPrefix(msg, "to|") {
		userName := strings.Split(msg, "|")[1]
		content := strings.Split(msg, "|")[2]
		remoteUser, ok := t.server.OnlineMap[userName]
		if ok {
			remoteUser.SendMessage(t.Name +"对您说："+content + "\n")
		} else {
			t.SendMessage("用户不存在")
			return
		}

	}else {
		t.server.BroadCast(t, msg)
	}

}
