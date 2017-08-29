package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	flag.Parse()
	workdir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(workdir)
}
