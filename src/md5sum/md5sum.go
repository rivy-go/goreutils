package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	flag.Parse()
	file, _ := ioutil.ReadFile(flag.Arg(0))
	hash := md5.Sum(file)
	hashString := fmt.Sprintf("%x", hash)
	fmt.Print(hashString)
}
