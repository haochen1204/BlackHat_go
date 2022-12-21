package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/miekg/dns"
)

func parse(filename string) (map[string]string, error) {
	records := make(map[string]string)
	fh, err := os.Open(filename)
	if err != nil {
		return records, err
	}
	defer fh.Close()
	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, ",", 2)
		if len(parts) < 2 {
			return records, fmt.Errorf("%s is not a valid line", line)
		}
		records[parts[0]] = parts[1]
	}
	return records, scanner.Err()
}

func main() {
	records, err := parse("proxy.config")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", records)
	dns.HandleFunc(".", func(w dns.ResponseWriter, req *dns.Msg) {
		if len(req.Question) < 1 {
			dns.HandleFailed(w, req)
			return
		}
		name := req.Question[0].Name
		parts := strings.Split(name, ".")
		if len(parts) > 1 {
			name = fmt.Sprintf("%s.%s", parts[len(parts)-3], parts[len(parts)-2])
		}
		match, ok := records[name]
		if !ok {
			dns.HandleFailed(w, req)
			return
		}
		fmt.Println(name)
		fmt.Println(match)
		fmt.Println(req)
		resp, err := dns.Exchange(req, match)
		if err != nil {
			fmt.Println(err)
			dns.HandleFailed(w, req)
			return
		}
		if err := w.WriteMsg(resp); err != nil {
			dns.HandleFailed(w, req)
			fmt.Println("1")
			return
		}
		fmt.Println(resp)
	})
	log.Fatal(dns.ListenAndServe(":53", "udp", nil))
}
