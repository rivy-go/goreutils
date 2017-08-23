package main

import (
	"crypto/sha512"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	flag.Parse()
	file, err := ioutil.ReadFile(flag.Arg(0))
	hash := sha512.Sum512(file)
	hashString := fmt.Sprintf("%x", hash)
	fmt.Print(hashString)
	if err != nil {
		fmt.Print("Error")
		os.Exit(1)
	}
}
