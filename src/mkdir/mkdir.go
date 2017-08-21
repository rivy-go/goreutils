package main

import (
	"flag"
	"os"
)

func main() {
	var mode int
	flag.IntVar(&mode, "mode", 775, "The file mode")
	flag.Parse()
	os.Mkdir(flag.Arg(0), os.FileMode(mode))
}
