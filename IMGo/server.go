package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)


type Server struct {
	Ip string
	Port int
	OnlineMap map[string]*User
	Message chan string
	mapLock sync.RWMutex
}

func NewServer (ip string, port int) *Server {
	server := &Server{
		Ip: ip,
		Port: port,
		OnlineMap: make(map[string]*User),
		Message: make(chan string),
	}


	return server
}

func (server *Server) ListenMessage () {
	for {
		msg := <- server.Message

		server.mapLock.Lock()

		for _, cli := range server.OnlineMap {
			cli.C <- msg
		}
		server.mapLock.Unlock()
	}
}

func (server *Server) Start () {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", server.Ip, server.Port))

	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}

	defer listener.Close()

	go server.ListenMessage()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener accept err:", err)
			continue
		}
		go server.Handler(conn)
	}
}

func (server *Server) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg
	server.Message <- sendMsg
}

func (server *Server) Handler (conn net.Conn) {
	
	user := NewUser(conn, server)

	user.Online()
	isLive := make(chan bool)

	go func () {
		buf := make([]byte, 40960)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				user.Offline()
				return
			}
			
			if err != nil && err != io.EOF {
				fmt.Println("Conn Read err:", err)
				return
			}

			msg := string(buf[:n-1])
			user.DoMessage(msg)
			isLive <- true
		}
	}()

	// WHY?
	for {
		select {
		case <- isLive:
		case <- time.After(time.Second * 100):
			user.SendMsg("你被踢了")
			close(user.C)
			conn.Close()
			return
		}
	}
}