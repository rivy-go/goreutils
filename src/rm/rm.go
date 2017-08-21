package main

import (
	"flag"
	"os"
)

func main() {
	flag.Parse()
	os.RemoveAll(flag.Arg(0))
}
