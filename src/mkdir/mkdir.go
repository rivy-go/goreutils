package main

import (
	"flag"
	"os"
)

func main() {
	var mode int
	flag.IntVar(&mode, "mode", 775, "The file mode")
	flag.Parse()
	os.Mkdir(os.Args[1], os.FileMode(mode))
}
