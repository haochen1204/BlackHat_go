package main

import (
	"net/http"
	"net/url"
)

func main() {
	form := url.Values{}
	form.Add("foo", "bar")
	r3, _ := http.PostForm("https://www.haochen1204.com", form)
	defer r3.Body.Close()
}
