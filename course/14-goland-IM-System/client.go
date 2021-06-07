package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

type Client struct {
	Ip   string
	Port int
	Name string
	Conn net.Conn
	flag int
}

func NewClient(serverIp string, serverPort int) *Client {
	// 创建客户端对象
	client := &Client{
		Ip:   serverIp,
		Port: serverPort,
		flag: 999,
	}
	// 链接server
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIp, serverPort))
	if err != nil {
		fmt.Println("net Dial")
		return nil
	}
	client.Conn = conn

	// 返回对象
	return client

}

var serverIp string
var serverPort int

//./client -ip 127.0.0.1 -port 8888
func init() {
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "设置服务器IP地址（默认127.0.0.1）")
	flag.IntVar(&serverPort, "port", 8000, "设置服务器Port（默认8000）")
}

func (t *Client) menu() bool {
	var flag int
	fmt.Println("1.公聊模式")
	fmt.Println("2.私聊模式")
	fmt.Println("3.更新用户名")
	fmt.Println("0.退出")
	fmt.Scanln(&flag)
	if flag >= 0 && flag <= 3 {
		t.flag = flag
		return true
	} else {
		fmt.Println("输入合法数字")
		return false
	}

}

func (t *Client) DealResponse() {
	// 一旦有数据旧直接copy 到 stdout标准输出上，永久阻塞监听
	io.Copy(os.Stdout, t.Conn)
	// 和上面的相等
	//for {
	//	buf := make([]byte, 1024)
	//	t.Conn.Read(buf)
	//}
}

func (t *Client) UpdateName() bool {
	fmt.Println("》》》》》 请输入用户名")
	fmt.Scanln(&t.Name)
	sendMsg := fmt.Sprintf("rename|%s\n", t.Name)
	_, err := t.Conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println(err)
		return false
	} else {
		return true
	}

}

func (t *Client) PublicChat() {
	var content string
	fmt.Println(">>>>> 请输入信息,exit 表示退出")
	fmt.Scanln(&content)
	for content != "exit" {
		sendMsg := fmt.Sprintf("%s\n", content)
		_, err := t.Conn.Write([]byte(sendMsg + "\n"))
		if err != nil {
			break
		}
		content = ""
		fmt.Println(">>>>> 请输入信息,exit 表示退出")
		fmt.Scanln(&content)
	}

}
func (t *Client) SelectUsers() {
	sendMsg := "who\n"
	t.Conn.Write([]byte(sendMsg))
}
func (t *Client) PrivateChat() {
	// 查询当前有哪些用户列表
	t.SelectUsers()
	var remoteName string
	fmt.Println(">>>>> 请输入聊天对象,exit 表示退出")
	fmt.Scanln(&remoteName)

	var content string
	fmt.Println(">>>>> 请输入聊天对象,exit 表示退出")
	fmt.Scanln(&content)

	for remoteName != "exit" {
		sendMsg := fmt.Sprintf("to|%s|%s\n", remoteName, content)
		_, err := t.Conn.Write([]byte(sendMsg + "\n"))
		if err != nil {
			break
		}
		content = ""
		fmt.Println(">>>>> 请输入信息,exit 表示退出")
		fmt.Scanln(&content)
	}

}
func (t *Client) Run() {
	for t.flag != 0 {
		for t.menu() != true {
		}
		switch t.flag {
		case 1: // 公聊模式
			fmt.Println("公聊模式")
			t.PublicChat()
			break
			//t.Conn.Write([]byte("dsds"))
		case 2: // 私聊模式
			fmt.Println("私聊模式")
			//t.Conn.Write([]byte("to"))
		case 3: // rename 模式
			fmt.Println("rename")
			t.UpdateName()
			//t.Conn.Write([]byte("to"))

		}
	}

}

func main() {
	//fmt.Println("请输入服务器ip")
	//fmt.Scan(&serverIp)
	//fmt.Println("请输入端口")
	//fmt.Scan(&serverPort)
	flag.Parse()
	client := NewClient(serverIp, serverPort)
	if client == nil {
		fmt.Println(">>>>> 链接服务器失败")
		return
	}
	go client.DealResponse()
	fmt.Println("》》》》》 链接服务器成功")
	//client.Conn.Write([]byte("hello\n"))

	// 启动客户端的业务
	//select {}
	client.Run()
}
