package main

import (
	"crypto/sha512"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	file, _ := ioutil.ReadFile(os.Args[1])
	hash := sha512.Sum512(file)
	hashString := fmt.Sprintf("%x", hash)
	fmt.Print(hashString)
}
