package main

import (
	"flag"
	"log"
	"os"
)

func main() {
	var mode int
	flag.IntVar(&mode, "mode", 700, "The file mode")
	flag.Parse()
	err := os.Mkdir(flag.Arg(0), os.FileMode(mode))
	if err != nil {
		log.Fatal(err)
	}
}
