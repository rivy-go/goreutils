package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Parse()
	workdir, err := os.Getwd()
	fmt.Print(workdir)
	if err != nil {
		fmt.Print("Error")
		os.Exit(1)
	}
}
