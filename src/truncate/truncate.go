package main

import (
	"flag"
	"log"
	"os"
)

func main() {
	flag.Parse()
	_, err := os.Stat(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}
	os.Truncate(flag.Arg(0), 0)
}
