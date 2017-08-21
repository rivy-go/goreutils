package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	file, _ := ioutil.ReadFile(os.Args[1])
	hash := md5.Sum(file)
	hashString := fmt.Sprintf("%x", hash)
	fmt.Print(hashString)
}
