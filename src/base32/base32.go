package main

import (
	"encoding/base32"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	var decode bool
	flag.BoolVar(&decode, "decode", false, "Decode the data instead of encode")
	flag.Parse()
	file, err := ioutil.ReadFile(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}
	var data string
	if decode {
		dataB, err := base32.StdEncoding.DecodeString(string(file))
		if err != nil {
			log.Fatal(err)
		}
		data = string(dataB)
	} else {
		data = base32.StdEncoding.EncodeToString(file)
	}
	fmt.Println(data)
}
