package main

import (
	"flag"
	"os"
)

func main() {
	flag.Parse()
	os.Rename(flag.Arg(0), flag.Arg(1))
}
