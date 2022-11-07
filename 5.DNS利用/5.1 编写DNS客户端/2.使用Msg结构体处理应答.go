package main

import (
	"fmt"

	"github.com/miekg/dns"
)

/*
type Msg struct {
	Msghdr
	Compress bool       `json:"-"` // 如果消息为true，消息将被压缩
	Question []Question // 保留question部分的RR
	Answer   []RR       // 保留answer部分的RR
	Ns       []RR       // 保留authority部分的RR
	Extra    []RR       // 保留addittional部分的RR
}
*/

func main() {
	var msg dns.Msg
	fqdn := dns.Fqdn("image.haochen1204.com")
	msg.SetQuestion(fqdn, dns.TypeA)
	in, err := dns.Exchange(&msg, "8.8.8.8:53")
	if err != nil {
		panic(err)
	}
	if len(in.Answer) < 1 {
		fmt.Println("No records")
		return
	}
	for _, answer := range in.Answer {
		if a, ok := answer.(*dns.A); ok {
			fmt.Println(a.A)
		}
	}
}
