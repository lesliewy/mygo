package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	test1()
}

func test1() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		//handlerConn(conn)   // 一次只处理一个连接.
		go handlerConn(conn)
	}
}

func handlerConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("2006-01-02 15:04:01\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
