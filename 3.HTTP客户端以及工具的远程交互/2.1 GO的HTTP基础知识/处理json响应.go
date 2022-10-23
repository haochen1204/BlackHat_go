package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Status struct {
	Message string
	Status  string
}

func main() {
	res, err := http.Post(
		"http://ip:port/ping",
		"application/json",
		nil,
	)
	if err != nil {
		log.Fatalln(err)
	}
	var status Status
	if err := json.NewDecoder(res.Body).Decode(&status); err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()
	log.Panicf("%s -> %s\n", status.Status, status.Message)
}
