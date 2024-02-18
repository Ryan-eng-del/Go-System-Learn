package block_net

import (
	"log"
	"net"
	"sync"
	"time"
)


func BlockIONet() {
	addr := "127.0.0.1:5678"

	wg := sync.WaitGroup{}

	wg.Add(2)

	go func (wg *sync.WaitGroup){
		defer wg.Done()
		conn, _ := net.Dial("tcp", addr)
		defer conn.Close()
		buf := make([]byte, 1024)
		log.Println("start read", time.Now().Format("03:04:05.000"))
		n, _ := conn.Read(buf)
		log.Println("content", string(buf[:n]),time.Now().Format("03:04:05.000"))
	}(&wg)


	go func (wg *sync.WaitGroup) {
		defer wg.Done()

		l, _ := net.Listen("tcp", addr)
		defer l.Close()
		for {
			conn, _ := l.Accept()
			go func(conn net.Conn) {
				defer conn.Close()
				log.Println("connected")
				time.Sleep(3*time.Second)
				_, err := conn.Write([]byte("Block I/O"))

				if err != nil {
					log.Println("error")
				}
			}(conn)
		}

	}(&wg)
	wg.Wait()
}