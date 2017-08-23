package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {
	flag.Parse()
	intMode, _ := strconv.Atoi(flag.Arg(0))
	mode := os.FileMode(intMode)
	err := os.Chmod(flag.Arg(1), mode)
	if err != nil {
		fmt.Print("Error")
		os.Exit(1)
	}
}
