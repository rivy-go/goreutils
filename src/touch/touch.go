package main

import (
	"flag"
	"os"
	"time"
)

func main() {
	var noCreate bool
	var mode int
	flag.BoolVar(&noCreate, "no-create", false, "Don't create new files")
	flag.IntVar(&mode, "mode", 666, "The filemode if it will be created")
	flag.Parse()
	if !noCreate {
		os.OpenFile(flag.Arg(0), os.O_CREATE, os.FileMode(mode))
	}
	currentTime := time.Now()
	os.Chtimes(flag.Arg(0), currentTime, currentTime)
}
