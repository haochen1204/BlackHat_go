package main

import (
	"io"
	"log"
	"net"
)

// echo 是一个处理函数，它仅回显接收到的数据
func echo(conn net.Conn) {
	defer conn.Close()
	// 创建一个缓冲区来存放接受到的数据
	b := make([]byte, 512)
	for {
		// 通过conn.Read接收数据到一个缓冲区
		size, err := conn.Read(b[0:])
		if err == io.EOF {
			log.Panicln("Client disconnected")
			break
		}
		if err != nil {
			log.Panicln("Unexpected error")
			break
		}
		log.Printf("Received %d bytes: %s\n", size, string(b))
		log.Println("Writing data")
		if _, err := conn.Write(b[0:size]); err != nil {
			log.Fatalln("Unable to write data")
		}
	}
}

func main() {
	// 在所有接口上绑定TCP端口20080
	listener, err := net.Listen("tcp", ":20080")
	if err != nil {
		log.Fatalln("unable to bind to port")
	}
	log.Println("Listening on 0.0.0.0:20080")
	for {
		// 等待连接，在已建立的连接上创建net.Conn
		conn, err := listener.Accept()
		log.Println("Received connection")
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}
		// 处理连接。使用goroutine实现并发
		go echo(conn)
	}
}
