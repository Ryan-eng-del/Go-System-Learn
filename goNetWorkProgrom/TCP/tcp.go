package tcp

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"sync"
	"time"
)

const TCP = "tcp"
var host, port = "localhost", 8120
var address = fmt.Sprintf("%s:%d", host, port)


type Message struct {
	ID uint `json:"id"`
	Content string `json:"content"`
	Code string `json:"code"`
	Time time.Time `json:"time"`
}


func TcpServer() {
	listener, err := net.Listen(TCP, address)
	if err != nil {
		log.Fatalln(err)
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
		}
		go HandleConn(conn)
	}
}


func HandleConn(conn net.Conn) {
	defer func () {
		conn.Close()
		log.Println("Server closed connection")
	}()
	wg := sync.WaitGroup{}
	wg.Add(1)
	// go ServerWrite(conn, &wg)
	go SendPingMessage(conn, &wg)
	// go ServerRead(conn, &wg)
	wg.Wait()
}


func SendPingMessage(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	curCount := 1
	maxCount := 3
	ctx, cancel := context.WithCancel(context.Background())
	go ReceivePongMessage(conn, ctx)

	ticker := time.NewTicker(2 * time.Second)
	for t := range ticker.C {
		pingMsg := Message{
			ID: uint(rand.Int()),
			Content: "Ping",
			Code: "Ping",
			Time: t,
		}

		encoder := json.NewEncoder(conn)
		if error := encoder.Encode(pingMsg); error != nil {
			log.Printf("Error encoding count is  %d", curCount)
			curCount += 1;
			if (maxCount < curCount) {
				cancel()
				log.Println("cancel ping request")
				return
			}
		}
		log.Println("Pint sent to server", conn.RemoteAddr())
	}
}

func ReceivePongMessage(conn net.Conn, ctx context.Context) {
	for {
		select {
		case  <- ctx.Done():
			return
		default:
			var message Message
			decoder := json.NewDecoder(conn)
			err := decoder.Decode(&message); 
			// 短链接,服务器已经关闭
			if err != nil && errors.Is(err, io.EOF) {
				fmt.Println("server close connection")
				return
			}
	
			if message.Code == "Pong" {
				log.Printf("from: %s, message-id: %d message-content: %s-%s\n", conn.RemoteAddr(),message.ID, message.Content, message.Code)
			}
		}
	}
}

func ServerWrite(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	var message Message = Message{
		ID: 1,
		Content: "Server Message",
		Code: "200",
	}
	encoder := json.NewEncoder(conn)
	if error := encoder.Encode(message); error != nil {
		fmt.Printf("Error encoding")
	}
	// wn, err := conn.Write([]byte("Hello I'm from Server" + "\n"))
	// if err != nil {
	// 	log.Println(err)
	// }
	// log.Printf("send data from server len is %d\n", wn)
}

func ServerRead(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	buf := make([]byte, 1024)
	rn, err := conn.Read(buf)

	if err != nil {
		log.Println(err)
	}
	log.Printf("received %d bytes from client", rn)
	log.Println(string(buf))
}



func TcpClient() {
	num := 1

	wg := sync.WaitGroup{}
	wg.Add(num)

	for i := 0; i < num; i++ {
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			
			conn, err := net.Dial(TCP, address)
			
			if err != nil {
				log.Fatal(err)
			}
			defer conn.Close()
			clientRwWg := sync.WaitGroup{}
			clientRwWg.Add(1)
			go ClientRead(conn,&clientRwWg)
			// go ClientWrite(conn, &clientRwWg)
			clientRwWg.Wait()
		}(&wg)
	}
	wg.Wait()
}


func ClientWrite(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	wn, err := conn.Write([]byte("Hello I'm from client" + "\n"))
	if err != nil {
		log.Println(err)
	}
	log.Printf("send data from client len is %d\n", wn)
}


func ClientRead(conn net.Conn,wg *sync.WaitGroup) {
	defer wg.Done()
	var message Message

	for {
		decoder := json.NewDecoder(conn)
		err := decoder.Decode(&message); 
		// 短链接,服务器已经关闭
		if err != nil && errors.Is(err, io.EOF) {
			fmt.Println("server close connection")
			break
		}

		if message.Code == "Ping" {
			// log.Println("send pont to server", conn.RemoteAddr())
			PongMessage(conn, &message)
		}

		// fmt.Printf("Message: %v\n", message)
	}

	// buf := make([]byte, 1024)
	// rn, err := conn.Read(buf)

	// if err != nil {
	// 	log.Println(err)
	// }
	// log.Printf("received %d bytes from server", rn)
	// log.Println(string(buf[:rn]))
}


func PongMessage(conn net.Conn, msg *Message) {
		pingMsg := Message{
			ID: uint(rand.Int()),
			Content: fmt.Sprintf("PingID(%d)", msg.ID),
			Code: "Pong",
			Time: time.Now(),
		}

		encoder := json.NewEncoder(conn)
		if error := encoder.Encode(pingMsg); error != nil {
			fmt.Printf("Error encoding")
		}
		log.Println("pong sent to server", conn.RemoteAddr())
}

