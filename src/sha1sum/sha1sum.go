package main

import (
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	file, _ := ioutil.ReadFile(os.Args[1])
	hash := sha1.Sum(file)
	hashString := fmt.Sprintf("%x", hash)
	fmt.Print(hashString)
}
