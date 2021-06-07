package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

type Server struct {
	Ip   string
	Port int
	// 在线用户类标
	OnlineMap map[string]*User
	mapLock   sync.RWMutex
	// 消息广播的channel
	Message chan string
}

// 创建一个server的接口

func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}
	return server
}

// 广播消息的方法

func (t *Server) BroadCast(user *User, msg string, ) {
	sendMsg := fmt.Sprintf("[%s]%s:%s", user.Addr, user.Name, msg)
	t.Message <- sendMsg

}

// 监听Message 广播消息channel 的goroutine 一旦有消息就发送给全部的在线User

func (t *Server) ListenMessage() {
	for {
		msg := <-t.Message
		// 将Msg 发送给全部的在线的用户
		t.mapLock.Lock()
		for _, cli := range t.OnlineMap {
			cli.C <- msg
		}
		t.mapLock.Unlock()

	}
}

func (t *Server) Handel(conn net.Conn) {
	// ....当前链接的业务
	fmt.Println("建立连接成功")
	// 用户上线,将用户加入到onlinemap 中
	user := NewUser(conn, t)
	user.Online()
	isLive := make(chan bool)
	// 接收客户端发送的消息， 然后广播
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				user.Offline()
				return
			}
			if err != nil && err != io.EOF {
				fmt.Println(err)
			}
			msg := string(buf[:n-1])
			user.DoMessage(msg)
			// 用户的任意消息，代表当前用户是一个活跃的
			isLive <- true

		}
	}()

	// 当前handler 阻塞
	for {
		select {
		// case 第一个满足会尝试走下面的，但后面的条件会刷新， 每次都有消息 都会走 isalive，然后刷新时间
		case <-isLive:
			// 当前用户是活跃的，应该重制定时器
			// 不做任何事情，为了激活select，更新下面的定时器
		case <-time.After(time.Second * 1000):
			// 已经超时，将当前的User 强制关闭
			user.SendMessage("你被提出")
			// 销毁调用资源
			close(user.C)
			// 关闭链接
			conn.Close()

		}
	}
}

// 启动服务器端口

func (t *Server) Start() {
	// socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", t.Ip, t.Port))
	if err != nil {
		fmt.Println("net.Listen err", err)
		return
	}
	// 启动监听Message 的 goroutine
	go t.ListenMessage()
	defer listener.Close()
	for {
		// accept
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listen accept err:", err)
			continue
		}
		// do handler
		go t.Handel(conn)
		// close listen socket
	}

}
