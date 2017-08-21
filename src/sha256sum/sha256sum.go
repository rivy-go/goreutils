package main

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	file, _ := ioutil.ReadFile(os.Args[1])
	hash := sha256.Sum256(file)
	hashString := fmt.Sprintf("%x", hash)
	fmt.Print(hashString)
}
