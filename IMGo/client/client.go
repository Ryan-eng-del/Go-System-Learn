package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

var serverIp string
var serverPort int


type Client struct {
	ServerIp string
	ServerPort int
	Name string
	Conn net.Conn
	flag int
}

func (client *Client) menu () bool {
	var flag int

	fmt.Println("1.公聊模式")
	fmt.Println("2.私聊模式")
	fmt.Println("3.更新用户名")
	fmt.Println("0.退出")

	fmt.Scanln(&flag)

	if flag >= 0 && flag <= 3 {
		client.flag = flag
		return true
	} else {
		fmt.Println(">>>>请输入合法范围内的数字<<<<")
		return false
	}
}
func (c *Client) updateName() bool{
	fmt.Println(">>>>请输入用户名:")
	fmt.Scanln(&c.Name)

	sendMsg := "rename|" + c.Name + "\n"

	_, err := c.Conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn.Write err:", err)
		return false
	}
	return true
}

func (c *Client) publicChat() {
	var chatMsg string

	
	fmt.Println(">>>>请输入聊天内容，exit退出.")
	fmt.Scanln(&chatMsg)

	for chatMsg != "exit" {
		if len(chatMsg) != 0 {
			sendMsg := chatMsg + "\n"
			_, err := c.Conn.Write([]byte(sendMsg))
			if err != nil {
				fmt.Println("conn Write err:", err)
				break
			}
		}

		chatMsg = ""
		fmt.Println(">>>>请输入聊天内容，exit退出.")
		fmt.Scanln(&chatMsg)
	}
}

//查询在线用户
func (client *Client) SelectUsers() {
	sendMsg := "who\n"
	_, err := client.Conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn Write err:", err)
		return
	}
}

func (c *Client) privateChat() {
	var remoteName string
	var chatMsg string
	c.SelectUsers()

	fmt.Println(">>>>请输入聊天对象[用户名], exit退出:")
	fmt.Scanln(&remoteName)

	for remoteName != "exit" {
		
		for chatMsg != "exit" {
			if len(chatMsg) != 0 {
				sendMsg := "to|" + remoteName + "|" + chatMsg + "\n\n"
				_, err := c.Conn.Write([]byte(sendMsg))
				if err != nil {
					fmt.Println("conn Write err:", err)
					break
				}
			}

			chatMsg = ""
			fmt.Println(">>>>请输入消息内容, exit退出:")
			fmt.Scanln(&chatMsg)
		}

		c.SelectUsers()
		fmt.Println(">>>>请输入聊天对象[用户名], exit退出:")
		fmt.Scanln(&remoteName)
	}
}

func (c *Client) Run() {
	for c.flag != 0 {
		for !c.menu() {
		}
		switch c.flag {
		case 1:
			c.publicChat()
		case 2:
			c.privateChat()
		case 3:
			c.updateName()
		}
	

	}

	
}
func NewClient(serverIp string, serverPort int) *Client {
	client := &Client{
		ServerIp: serverIp,
		ServerPort: serverPort,
		flag: 999,
	}

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIp, serverPort))

	if err != nil {
		fmt.Println("net.Dial error:", err)
		return nil
	}

	client.Conn = conn
	return client
}
func init() {
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "设置 ip 地址，默认是 127.0.0.1")
	flag.IntVar(&serverPort, "port", 8888, "设置 port，默认是 8888")
}

func (c *Client) DealResponse () {
	io.Copy(os.Stdout, c.Conn)
}

func main() {
	flag.Parse()

	client := NewClient(serverIp, serverPort)
	if client == nil {
		fmt.Println(">>>>> 链接服务器失败...")
		return
	}

	fmt.Println(">>>>>链接服务器成功...")
	go client.DealResponse()
	client.Run()

	
}