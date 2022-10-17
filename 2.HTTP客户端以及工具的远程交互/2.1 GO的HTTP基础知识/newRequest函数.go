package main

import (
	"net/http"
	"net/url"
	"strings"
)

func main() {
	form := url.Values{}
	form.Add("foo", "bar")
	var client http.Client
	req, _ := http.NewRequest(
		"POST",
		"https://www.haochen1204.com",
		strings.NewReader(form.Encode()),
	)
	resp, _ := client.Do(req)
	defer resp.Body.Close()
}
