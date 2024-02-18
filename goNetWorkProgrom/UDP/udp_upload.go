package udp

import (
	"errors"
	"io"
	"log"
	"net"
	"os"
)


func UDPFileClient()  {
	filename := "./data/sample-6s.mp3"
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}

	addr := "127.0.0.1:6789"
	raddr, err := net.ResolveUDPAddr("udp", addr)

	if err != nil {
		log.Fatal(err)
	}

	udpConn, err := net.DialUDP("udp", nil, raddr)

	if err != nil {
		log.Fatal(err)
	}

	defer udpConn.Close()
	

	if _, err := udpConn.Write([]byte(fileInfo.Name())); err != nil {
		log.Fatal(err)
	}
	log.Println("文件大小", fileInfo.Size())

	buf := make([]byte, 4*1024)
	rn, err := udpConn.Read(buf)
	if err != nil {
		log.Fatal(err)
	}

	if string(buf[:rn]) != "ok"{
		log.Fatal(errors.New("server not ready"))
	}

	i := 0

	for {
		rn, err := file.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
			break
		}
		if _, err := udpConn.Write(buf[:rn]); err != nil {
			log.Fatal(err)
		}
		i++
		log.Println("file write some content", i)
	}
}


func UDPFileServer () {
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:6789")

	if err != nil {
		log.Fatal(err)
	}

	log.Println(addr.String(), "addr")
	udpConn, err := net.ListenUDP("udp", addr)

	if err != nil {
		log.Fatal(err)
	}

	defer udpConn.Close()

	log.Printf("server listening on %s", udpConn.LocalAddr())


	buf := make([]byte, 4 * 1024)
	rn, raddr, err := udpConn.ReadFromUDP(buf);
	if err != nil {
		log.Fatal(err)
	}
	filename := buf[:rn]
	log.Println("filename: ",string(filename))
	log.Println("filename: ",raddr.String(), udpConn.RemoteAddr())
	if _, err := udpConn.WriteToUDP([]byte("ok"), raddr); err != nil {
		log.Fatal(err)
	}

	file, err := os.Create(string(filename))
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	for {
		rn, _, err:= udpConn.ReadFromUDP(buf)
		if err != nil {
			log.Fatalln(err)
		}

		if _, err := file.Write(buf[:rn]); err != nil {
			log.Fatal(err)
		}

		log.Println("server write some content")
	}
}