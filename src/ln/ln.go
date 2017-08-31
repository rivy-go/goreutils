package main

import (
	"flag"
	"log"
	"os"
)

func main() {
	var symbolic bool
	flag.BoolVar(&symbolic, "symbolic", false, "Specify if this is a hard or symbolic link.")
	flag.Parse()
	var err error
	if symbolic {
		err = os.Symlink(flag.Arg(0), flag.Arg(1))
	} else {
		err = os.Link(flag.Arg(0), flag.Arg(1))
	}
	if err != nil {
		log.Fatal(err)
	}
}
