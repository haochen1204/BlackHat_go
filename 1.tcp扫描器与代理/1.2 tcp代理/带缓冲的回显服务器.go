package main

import (
	"io"
	"log"
	"net"
)

// echo 是一个处理函数，它仅回显接收到的数据
func echo(conn net.Conn) {
	defer conn.Close()
	/*
		reader := bufio.NewReader(conn)
		s, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalln("Unable to read data")
		}
		log.Printf("Read %d bytes: %s ", len(s), s)
		log.Println("Writing data")
		writer := bufio.NewWriter(conn)
		if _, err := writer.WriteString(s); err != nil {
			log.Fatalln("Unable to write data")
		}
		writer.Flush()
	*/
	if _, err := io.Copy(conn, conn); err != nil {
		log.Fatalln("unable to read/write data")
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
