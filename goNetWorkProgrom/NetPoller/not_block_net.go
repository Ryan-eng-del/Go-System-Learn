package block_net

import (
	"log"
	"net"
	"sync"
	"time"
)


func NotBlockIONet() {
	addr := "127.0.0.1:5678"

	wg := sync.WaitGroup{}

	wg.Add(2)

	go func (wg *sync.WaitGroup){
		defer wg.Done()
		conn, _ := net.Dial("tcp", addr)
		defer conn.Close()
		buf := make([]byte, 1024)
		log.Println("start read", time.Now().Format("03:04:05.000"))
		conn.SetReadDeadline(time.Now().Add(400 *time.Millisecond))
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
				rn, err := conn.Write([]byte("Block I/O"))
				if err != nil {
					log.Println(err)
				}
				log.Printf("read content length is %d", rn)
			}(conn)
		}
	}(&wg)
	wg.Wait()
}

func NotBlockIOChannelNet() {
	addr := "127.0.0.1:5678"

	wg := sync.WaitGroup{}
	ch := make(chan string)

	wg.Add(2)

	go func (wg *sync.WaitGroup){
		defer wg.Done()
		conn, _ := net.Dial("tcp", addr)
		defer conn.Close()
		log.Println("start read", time.Now().Format("03:04:05.000"))
		conn.SetReadDeadline(time.Now().Add(400 *time.Millisecond))
		// n, _ := conn.Read(buf)
		content := ""
		select {
		case content = <-ch:
		default:
		}
		
		log.Println("content", content,time.Now().Format("03:04:05.000"))
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
				rn, err := conn.Write([]byte("Block I/O"))
				if err != nil {
					log.Println(err)
				}
				log.Printf("read content length is %d", rn)
			}(conn)
		}
	}(&wg)
	wg.Wait()
}

func NotBlockIOChannelGoRoutineNet() {
	addr := "127.0.0.1:5678"

	wg := sync.WaitGroup{}

	wg.Add(2)

	go func (wg *sync.WaitGroup){
		defer wg.Done()
		conn, _ := net.Dial("tcp", addr)
		defer conn.Close()
		log.Println("start read", time.Now().Format("03:04:05.000"))
		conn.SetReadDeadline(time.Now().Add(400 *time.Millisecond))
		// n, _ := conn.Read(buf)

		wgwg := sync.WaitGroup{}
		wgwg.Add(1)
		data := make(chan []byte, 1024)
		go func(wg *sync.WaitGroup){
			defer wg.Done()
			buf := make([]byte, 1024)
			n, _ := conn.Read(buf)
			data <- buf[:n]
			log.Println("read content is ", string(buf[:n]))
		}(&wgwg)


		time.Sleep(time.Second * 1)

		content := []byte("")

		select {
		case content = <-data:
		default:
		}
		wgwg.Wait()
	
		log.Println("content", string(content),time.Now().Format("03:04:05.000"))
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
				// time.Sleep(3*time.Second)
				rn, err := conn.Write([]byte("Block I/O"))
				if err != nil {
					log.Println(err)
				}
				log.Printf("write content length is %d", rn)
			}(conn)
		}
	}(&wg)
	wg.Wait()
}