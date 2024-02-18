package main

import (
	"net"
	"strings"
)


type User struct {
	Name string
	Addr string
	C chan string
	conn net.Conn
	Server *Server
}

func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()

	user := &User{
		Name: userAddr,
		Addr: userAddr,
		C: make(chan string),
		conn: conn,
		Server: server,
	}

	go user.ListenMessage()

	return user
}

func (user *User) ListenMessage () {
	for {
		msg := <- user.C
		user.conn.Write([]byte(msg + "\n"))
	}
}

func (user *User) Online() {
	user.Server.mapLock.Lock()
	user.Server.OnlineMap[user.Name] = user
	user.Server.mapLock.Unlock()
	user.Server.BroadCast(user, "已上线")
}

func (user *User) Offline() {
	user.Server.mapLock.Lock()
	delete(user.Server.OnlineMap, user.Name)
	user.Server.mapLock.Unlock()
	user.Server.BroadCast(user, "已下线")

}

func (user *User) SendMsg(msg string) {
	user.conn.Write([]byte(msg))
}

//用户处理消息的业务
func (user *User) DoMessage(msg string) {

	if msg == "who"  {
		user.Server.mapLock.Lock()
		for _, user_online := range user.Server.OnlineMap {
			onlineMsg := "[" + user_online.Addr + "]" + user_online.Name + ":" + "在线...\n"
			user.SendMsg(onlineMsg)
		}
		user.Server.mapLock.Unlock()
	} else if len(msg) > 7 && msg[:7] == "rename|" {
		newName := strings.Split(msg, "|")[1]

		_, ok := user.Server.OnlineMap[newName]

		if ok {
			user.SendMsg("当前用户名被使用\n")
		} else {
			user.Server.mapLock.Lock()
			delete(user.Server.OnlineMap, user.Name)
			user.Server.OnlineMap[newName] = user
			user.Server.mapLock.Unlock()

			user.Name = newName
			user.SendMsg("您已经更新用户名:" + user.Name + "\n")
		}
	} else if len(msg) > 4 && msg[:3] == "to|" {
		//消息格式:  to|张三|消息内容

		//1 获取对方的用户名
		remoteName := strings.Split(msg, "|")[1]
		if remoteName == "" {
			user.SendMsg("消息格式不正确，请使用 \"to|张三|你好啊\"格式。\n")
			return
		}

		//2 根据用户名 得到对方User对象
		remoteUser, ok := user.Server.OnlineMap[remoteName]
		if !ok {
			user.SendMsg("该用户名不不存在\n")
			return
		}

		//3 获取消息内容，通过对方的User对象将消息内容发送过去
		content := strings.Split(msg, "|")[2]
		if content == "" {
			user.SendMsg("无消息内容，请重发\n")
			return
		}
		remoteUser.SendMsg(user.Name + "对您说:" + content)

	}else {
		user.Server.BroadCast(user, msg)
	}
}