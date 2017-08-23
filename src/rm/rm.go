package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Parse()
	err := os.RemoveAll(flag.Arg(0))
	if err != nil {
		fmt.Print("Error")
		os.Exit(1)
	}
}
