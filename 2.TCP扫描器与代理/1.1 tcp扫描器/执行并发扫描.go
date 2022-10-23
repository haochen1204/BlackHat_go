package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {
	// 创建同步计数器
	var wg sync.WaitGroup
	for i := 1; i <= 1024; i++ {
		// 计数器加一
		wg.Add(1)
		go func(j int) {
			// 在函数执行结束时计数器减一
			defer wg.Done()
			address := fmt.Sprintf("www.haochen1204.com:%d", j)
			//fmt.Println(address)
			cnn, err := net.Dial("tcp", address)
			if err != nil {
				//fmt.Println("close")
				return
			}
			cnn.Close()
			fmt.Printf("%d open\n", j)
		}(i)
	}
	// 阻塞 等待计数器归零
	wg.Wait()
}
