package main

import (
	"encoding/base32"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	flag.Parse()
	file, err := ioutil.ReadFile(flag.Arg(0))
	hash := base32.StdEncoding.EncodeToString(file)
	fmt.Println(hash)
	if err != nil {
		log.Fatal(err)
	}
}
