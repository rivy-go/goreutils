package main

import (
	"flag"
	"log"
	"os"
)

func main() {
	flag.Parse()
	err := os.RemoveAll(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}
}
