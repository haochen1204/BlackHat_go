package main

import (
	"net/http"
	"net/url"
	"strings"
)

func main() {
	r1, _ := http.Get("https://www.haochen1204.com")
	// 读取相应正文，未显示
	defer r1.Body.Close()
	r2, _ := http.Head("https://www.haochen1204.com")
	// 读取相应正文，未显示
	defer r2.Body.Close()
	form := url.Values{}
	form.Add("foo", "bar")
	r3, _ := http.Post(
		"https://www.haochen1204.com",
		"application/x-www-form-urlencoded",
		strings.NewReader(form.Encode()),
	)
	// 读取相应正文，未显示
	defer r3.Body.Close()
}
