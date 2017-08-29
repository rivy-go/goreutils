package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	flag.Parse()
	file, err := ioutil.ReadFile(flag.Arg(0))
	fmt.Print(string(file))
	if err != nil {
		log.Fatal(err)
	}
}
