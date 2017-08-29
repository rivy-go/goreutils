package main

import (
	"flag"
	"log"
	"os"
	"strconv"
)

func main() {
	flag.Parse()
	intMode, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}
	mode := os.FileMode(intMode)
	err = os.Chmod(flag.Arg(1), mode)
	if err != nil {
		log.Fatal(err)
	}
}
