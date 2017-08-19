package main

import (
	"crypto/sha512"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	file, _ := ioutil.ReadFile(os.Args[1])
	fmt.Printf("%x", sha512.Sum512(file))
}
