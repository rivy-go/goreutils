package main

import (
	"flag"
	"os"
	"time"
)

func main() {
	var noCreate bool
	flag.BoolVar(&noCreate, "no-create", false, "Don't create new files")
	flag.Parse()
	if noCreate == true {
		os.OpenFile(os.Args[1], os.O_CREATE, 0666)
	}
	currentTime := time.Now()
	os.Chtimes(os.Args[1], currentTime, currentTime)
}
