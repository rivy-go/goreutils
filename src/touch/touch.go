package main

import (
	"os"
	"time"
)

func main() {
	os.OpenFile(os.Args[1], os.O_CREATE, 0666)
	currentTime := time.Now()
	os.Chtimes(os.Args[1], currentTime, currentTime)
}
