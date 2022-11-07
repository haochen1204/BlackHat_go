package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var (
		flDomain      = flag.String("Domain", "", "The domain to perform guessing against.")
		flWordlist    = flag.String("Wordlist", "", "The wordlist to use for guessing.")
		flWorkerCount = flag.Int("c", 100, "The amount of workers to use.")
		flServerAddr  = flag.String("server", "8.8.8.8:53", "The DNS server to use.")
	)
	flag.Parse()

	if *flDomain == "" || *flWordlist == "" {
		fmt.Println("-domain and -wordlist are required")
		os.Exit(1)
	}
	fmt.Println(*flWorkerCount, *flServerAddr)
}
