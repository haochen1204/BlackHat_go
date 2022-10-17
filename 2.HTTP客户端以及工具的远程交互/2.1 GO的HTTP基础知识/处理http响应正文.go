package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://www.haochen1204.com")
	if err != nil {
		log.Panicln(err)
	}
	// 打印HTTP状态码
	fmt.Println(resp.Status)
	// 读取并显示响应正文
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println(string(body))
	resp.Body.Close()
}
