package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Parse()
	workdir, _ := os.Getwd()
	fmt.Print(workdir)
}
