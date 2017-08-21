package main

import (
	"crypto/sha1"
	"flag"
	"fmt"
	"io/ioutil"
)

func main() {
	flag.Parse()
	file, _ := ioutil.ReadFile(flag.Arg(0))
	hash := sha1.Sum(file)
	hashString := fmt.Sprintf("%x", hash)
	fmt.Print(hashString)
}
