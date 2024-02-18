package udp

import (
	"log"
	"net"
)


func UDPServer () {
	addr, err := net.ResolveUDPAddr("udp", ":9876")

	if err != nil {
		log.Fatalln(err)
	}

	udpConn, err := net.ListenUDP("udp", addr)

	if err != nil {
		log.Fatalln(err)
	}
	defer udpConn.Close()
	log.Printf("server is listening on %s", udpConn.LocalAddr().String())

	// 读
	buf := make([]byte, 1024)

	rn, addr, err := udpConn.ReadFromUDP(buf)

	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("received %s from %s", string(buf[:rn]), addr.String())

	// 写
	data := []byte("server " + string(buf[:rn]))
	wn, err := udpConn.WriteToUDP(data, addr)

	if err != nil {
		log.Fatalln(err, wn)
	}

	log.Printf("send %s to %s", string(data),addr.String() )

}


func UDPClient () {
	addr, err := net.ResolveUDPAddr("udp", ":9876")

	if err != nil {
		log.Fatalln(err)
	}
	udpConn, err := net.DialUDP("udp", nil, addr)

	if err != nil {
		log.Fatalln(err)
	}
	defer udpConn.Close()

	// 写
	data := []byte("from client")
	wn, err := udpConn.Write(data)

	if err != nil {
		log.Fatalln(err, wn)
	}

	log.Printf("send %s to %s", string(data),addr.String() )

	// 读
	buf := make([]byte, 1024)

	rn, addr, err := udpConn.ReadFromUDP(buf)

	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("received %s from %s", string(buf[:rn]), addr.String())


}