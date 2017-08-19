package main

import (
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	file, _ := ioutil.ReadFile(os.Args[1])
	fmt.Printf("%x", sha1.Sum(file))
}
