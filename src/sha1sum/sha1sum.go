package main

import (
	"crypto/sha1"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	flag.Parse()
	file, err := ioutil.ReadFile(flag.Arg(0))
	hash := sha1.Sum(file)
	hashString := fmt.Sprintf("%x", hash)
	fmt.Println(hashString)
	if err != nil {
		log.Fatal(err)
	}
}
