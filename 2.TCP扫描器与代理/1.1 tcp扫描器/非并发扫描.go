package main

import (
	"fmt"
	"net"
)

func main() {
	for i := 1; i <= 1024; i++ {
		address := fmt.Sprintf("www.haochen1204.com:%d", i)
		//fmt.Println(address)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			// 端口被关闭
			//fmt.Printf("%d close\n", i)
			continue
		}
		conn.Close()
		fmt.Printf("%d open\n", i)
	}
}
