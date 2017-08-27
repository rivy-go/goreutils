package main

import (
	"flag"
	"fmt"
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
		_, err := os.OpenFile(flag.Arg(0), os.O_CREATE, os.FileMode(mode))
		if err != nil {
			fmt.Print(err)
			os.Exit(1)
		}
	}
	currentTime := time.Now()
	err := os.Chtimes(flag.Arg(0), currentTime, currentTime)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}
