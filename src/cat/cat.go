package main

import (
	"flag"
	"fmt"
	"io/ioutil"
)

func main() {
	flag.Parse()
	file, _ := ioutil.ReadFile(flag.Arg(0))
	fmt.Print(string(file))
}
