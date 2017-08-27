package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Parse()
	err := os.Rename(flag.Arg(0), flag.Arg(1))
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}
