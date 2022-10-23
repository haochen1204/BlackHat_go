package main

import (
	"fmt"
	"log"
	"os"

	"github.com/haochen1204/msf"
)

func main() {
	host := os.Getenv("MSFHOST")
	pass := os.Getenv("MSFPASS")
	user := "msf"

	if host == "" || pass == "" {
		log.Fatalln("Missing required enviroment variable MSFHOST of MSFPASS")
	}
	msf, err := msf.New(host, user, pass)
	if err != nil {
		log.Panicln(err)
	}
	defer msf.Logout()
	sessions, err := msf.SessionList()
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println("Sessions:")
	for _, session := range sessions {
		fmt.Printf("%5d %s\n", session.ID, session.Info)
	}
}
