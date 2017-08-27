package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Parse()
	workdir, err := os.Getwd()
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	fmt.Print(workdir)
}
