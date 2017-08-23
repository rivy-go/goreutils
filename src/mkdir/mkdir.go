package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var mode int
	flag.IntVar(&mode, "mode", 775, "The file mode")
	flag.Parse()
	err := os.Mkdir(flag.Arg(0), os.FileMode(mode))
	if err != nil {
		fmt.Print("Error")
		os.Exit(1)
	}
}
