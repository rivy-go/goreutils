package main

import (
	"crypto/sha256"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	flag.Parse()
	file, err := ioutil.ReadFile(flag.Arg(0))
	hash := sha256.Sum256(file)
	hashString := fmt.Sprintf("%x", hash)
	fmt.Print(hashString)
	if err != nil {
		fmt.Print("Error")
		os.Exit(1)
	}
}
