package poll

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"sync"
	"time"
)

type Poll interface {
	Get() (net.Conn, error)
	Put(conn net.Conn) error
	Release() error
	Len() (int, error)
}



type PollConfig struct {
	InitConnection int
	MinConnection int
	MaxConnection int
	MaxIdle int
	IdleTimeout time.Duration
	Factory ConnFactory
	addr string
}


type IdleConn struct {
	Conn net.Conn
	PutTime time.Time
}


type TCPPoll struct {
	config *PollConfig
	openingNum int
	idleList chan *IdleConn
	mu *sync.RWMutex
}

var f, _= os.OpenFile("log.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)



func (poll *TCPPoll) Get() (net.Conn, error) {
	log.SetOutput(f)
	poll.mu.Lock()
	defer poll.mu.Unlock()

	for {
		select {
		case idleConn, ok := <- poll.idleList:
			if !ok {
				continue
			}         
			
			if idleConn.PutTime.Add(poll.config.IdleTimeout).Before(time.Now()) {
				log.Println("过期关闭")
				poll.config.Factory.Close(idleConn.Conn)
				continue
			}

			if err := poll.config.Factory.Ping(idleConn.Conn); err != nil {
				log.Println("ping 不通关闭")
				poll.config.Factory.Close(idleConn.Conn)
				continue
			}
			log.Println("from idle connection established")
			poll.openingNum++
			log.Printf("在线链接：%d, 链接池：%d",poll.openingNum, len(poll.idleList))
			return idleConn.Conn, nil
		default: 
			if poll.openingNum >= poll.config.MaxConnection {
				log.Printf("在线链接：%d, 链接池：%d",poll.openingNum, len(poll.idleList))
				log.Println("continue 阻塞")
				continue
			}

			conn, err := poll.config.Factory.Factory(poll.config.addr); 

			if err != nil {
				return nil,  err
			}
			poll.openingNum++
			log.Println("from factory connection established")
			log.Printf("在线链接：%d, 链接池：%d",poll.openingNum, len(poll.idleList))
			return conn, nil
		}
	}
}

func (poll *TCPPoll) Put(conn net.Conn, w *sync.WaitGroup) error {
	poll.mu.Lock()
	defer poll.mu.Unlock()
	defer w.Done()
	log.SetOutput(f)
	log.Println("come in Put method")
	if conn == nil {
		return errors.New("put connection error")
	}

	if poll.idleList == nil {
		log.Println("idleList is nil 关闭")
		poll.config.Factory.Close(conn)
		return errors.New("put connection error")
	}
	
	select {
	case poll.idleList <- &IdleConn{
		Conn: conn,
		PutTime: time.Now(),
	}: 
		poll.openingNum-=1
		log.Println("put connection established", len(poll.idleList))
		return nil
	default:
		poll.config.Factory.Close(conn)
		poll.openingNum-=1
		log.Println("put connection established", len(poll.idleList))
		return nil
	}
}

func (poll *TCPPoll) Release() error {
	poll.mu.Lock()
	defer poll.mu.Unlock()

	if poll.idleList == nil {
		return nil
	}

	close(poll.idleList)

	for conn := range poll.idleList {
		poll.config.Factory.Close(conn.Conn)
	}
	return nil
}

func (poll *TCPPoll) Len() (int, error) {
	panic("not implemented") // TODO: Implement
}

const DEFAULT_MAX_CONNECTIONS = 100
const DEFAULT_INIT_CONNECTIONS = 1

func NewTCPPoll (addr string, pollConfig PollConfig) (*TCPPoll, error) {
	if addr == ""  {
		return nil, errors.New("addr must not be empty")
	}

	if pollConfig.Factory == nil {
		return nil, errors.New("pollConfig.Factory must not be nil")
	}

	pollConfig.addr = addr

	if pollConfig.MaxConnection == 0 {
		pollConfig.MaxConnection = DEFAULT_MAX_CONNECTIONS
	}

	if pollConfig.IdleTimeout == 0 {
		pollConfig.IdleTimeout = time.Minute * 7
	}

	if pollConfig.InitConnection == 0 {
		pollConfig.InitConnection = DEFAULT_INIT_CONNECTIONS
	} else if pollConfig.InitConnection > pollConfig.MaxConnection {
		pollConfig.InitConnection = pollConfig.MaxConnection
	}

	if pollConfig.MaxIdle == 0 {
		pollConfig.MaxIdle = DEFAULT_INIT_CONNECTIONS
	} else if pollConfig.MaxIdle > pollConfig.MaxConnection {
		pollConfig.MaxIdle = pollConfig.MaxConnection
	}

	poll := TCPPoll {
		config: &pollConfig,
		openingNum: 0,
		idleList: make(chan *IdleConn, pollConfig.MaxConnection),
		mu: &sync.RWMutex{},
	}

	for i := 0; i < pollConfig.InitConnection; i++ {
		conn, err := poll.config.Factory.Factory(addr)

		if err != nil {
			log.Println("release connection 关闭")
			poll.Release()
			return nil, err
		}

		poll.idleList <- &IdleConn{
			Conn: conn,
			PutTime: time.Now(),
		}
	}

	return &poll, nil

	
}


type ConnFactory interface {
	Factory(addr string) (net.Conn, error)
	Close(net.Conn) error
	Ping(net.Conn) error
}

type TCPConnFactory struct {
}

func (mp *TCPConnFactory) Factory(addr string) (net.Conn, error) {
	if addr == ""  {
		return nil, errors.New("addr must not be empty")
	}
	
	conn, err := net.DialTimeout("tcp", addr, 100*time.Second)

	if err != nil {
		return nil ,err
	}
	return conn, nil

}

func (mp *TCPConnFactory) Close(conn net.Conn) error {
	return conn.Close()
}

func (mp *TCPConnFactory) Ping(_ net.Conn) error {
	return nil
}



// TCP Server
type Message struct {
	ID uint `json:"id"`
	Content string `json:"content"`
	Code string `json:"code"`
	Time time.Time `json:"time"`
}

func TcpServer() {
	listener, err := net.Listen("tcp", "127.0.0.1:5678")
	if err != nil {
		log.Fatalln(err)
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
		}
		log.Println("收到了", conn.RemoteAddr())
		go HandleConn(conn)
	}
}

func HandleConn(conn net.Conn) {
	// defer func () {
	// 	conn.Close()
	// 	log.Println("Server closed connection")
	// }()
	wg := sync.WaitGroup{}
	wg.Add(1)
	go ServerWrite(conn, &wg)
	// select
	wg.Wait()
}


func ServerWrite(conn net.Conn, wg *sync.WaitGroup) {
	// defer wg.Done()
	var message Message = Message{
		ID: 1,
		Content: "Server Message",
		Code: "200",
	}
	encoder := json.NewEncoder(conn)
	if error := encoder.Encode(message); error != nil {
		fmt.Printf("Error encoding")
	}
}