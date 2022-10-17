package main

import (
	"io"
	"log"
	"net"
)

func handler(src net.Conn) {
	dst, err := net.Dial("tcp", "www.haochen1204.com:443")
	if err != nil {
		log.Fatalln("Unable to connect to our unreachable host")
	}
	defer dst.Close()

	// 在goroutine中运行防止io.Copy被阻塞
	go func() {
		if _, err := io.Copy(dst, src); err != nil {
			log.Fatalln(err)
		}
	}()
	// 将目标的输出复制回源
	if _, err := io.Copy(src, dst); err != nil {
		log.Fatalln(err)
	}
}

func main() {
	// 在本地端口监听
	listener, err := net.Listen("tcp", ":1204")
	if err != nil {
		log.Fatalln("Unable to bind to port")
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}
		go handler(conn)
	}
}
