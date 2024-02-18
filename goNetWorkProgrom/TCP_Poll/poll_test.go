package poll

import (
	"log"
	"sync"
	"testing"
)

func TestPoll(t *testing.T) {
	serverAddress := "127.0.0.1:5678"

	poll, err := NewTCPPoll(serverAddress, PollConfig{
		Factory: &TCPConnFactory{},
		InitConnection: 2,
		MaxConnection: 10,
	})

	if err != nil {
		log.Fatal(err)
	}

	log.Println(poll.idleList, poll.openingNum, len(poll.idleList))

	wg := sync.WaitGroup{}
	clientNum := 300
	wg.Add(clientNum)

	for i := 0; i < clientNum; i++ {
		go func (wg *sync.WaitGroup, i int) {
				defer wg.Done()
				conn, err := poll.Get()
				if err != nil {
					log.Fatal(err)
					return
				}
				w := &sync.WaitGroup{}
				w.Add(1)
				go poll.Put(conn, w)
				w.Wait()
		}(&wg, i)
	}
	wg.Wait()
	poll.Release()
	log.Println("final", len(poll.idleList))
}


func TestTcpServer(t *testing.T) {
	TcpServer()
}

