package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	flag.Parse()
	content, err := ioutil.ReadFile(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.Create(flag.Arg(1))
	if err != nil {
		log.Fatal(err)
	}
	_, err = file.Write(content)
	if err != nil {
		log.Fatal(err)
	}
}
