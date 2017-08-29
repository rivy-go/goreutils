package main

import (
	"flag"
	"log"
	"os"
)

func main() {
	flag.Parse()
	err := os.Link(flag.Arg(0), flag.Arg(1))
	if err != nil {
		log.Fatal(err)
	}
}
