package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"io/ioutil"
)

func main() {
	flag.Parse()
	file, _ := ioutil.ReadFile(flag.Arg(0))
	hash := md5.Sum(file)
	hashString := fmt.Sprintf("%x", hash)
	fmt.Print(hashString)
}
