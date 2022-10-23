package main

import (
	"fmt"
	"net"
)

func main() {
	_, err := net.Dial("tcp", "www.haochen1204.com:80")
	if err == nil {
		fmt.Println("Connection successful!")
	} else {
		fmt.Println("Connection failed!")
	}
}
