package main

import (
	"flag"
	"log"
	"os"
)

func main() {
	flag.Parse()
	err := os.Rename(flag.Arg(0), flag.Arg(1))
	if err != nil {
		log.Fatal(err)
	}
}
