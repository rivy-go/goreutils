package main

import (
	"flag"
	"os"
	"strconv"
)

func main() {
	flag.Parse()
	intMode, _ := strconv.Atoi(flag.Arg(0))
	mode := os.FileMode(intMode)
	os.Chmod(flag.Arg(1), mode)

}
