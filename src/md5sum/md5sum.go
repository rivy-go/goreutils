package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	flag.Parse()
	file, err := ioutil.ReadFile(flag.Arg(0))
	hash := md5.Sum(file)
	hashString := fmt.Sprintf("%x", hash)
	fmt.Print(hashString)
	if err != nil {
		fmt.Print("Error")
		os.Exit(1)
	}
}
