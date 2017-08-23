package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	flag.Parse()
	file, err := ioutil.ReadFile(flag.Arg(0))
	fmt.Print(string(file))
	if err != nil {
		fmt.Print("Error")
		os.Exit(1)
	}
}
