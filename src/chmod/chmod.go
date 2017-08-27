package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {
	flag.Parse()
	intMode, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	mode := os.FileMode(intMode)
	err := os.Chmod(flag.Arg(1), mode)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}
