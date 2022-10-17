package main

import (
	"fmt"

	"github.com/haochen1204/go_hack"
)

func main() {
	s := go_hack.New_FoFa_Client("haochen1204@aliyun.com", "1664027a75442dde2fbd7f8825397186")
	ret, _ := s.APIInfo()
	fmt.Printf("email is %s and UserName is %s\n", ret.Email, ret.UserName)
	search_msg := go_hack.New_FoFa_InfoSearch("www.haochen1204.com")
	search_msg.Fields = "ip,port,host,country"
	msg, err := s.HostSearch(search_msg)
	if err != nil {
		fmt.Println("Fofa error ! ", err)
	} else {
		fmt.Println(msg.Results)
		for _, value := range msg.Results {
			fmt.Println(value[0], value[1], value[2], value[3])
		}
	}

}
