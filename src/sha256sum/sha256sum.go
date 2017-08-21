package main

import (
	"crypto/sha256"
	"flag"
	"fmt"
	"io/ioutil"
)

func main() {
	flag.Parse()
	file, _ := ioutil.ReadFile(flag.Arg(0))
	hash := sha256.Sum256(file)
	hashString := fmt.Sprintf("%x", hash)
	fmt.Print(hashString)
}
