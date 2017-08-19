package main

import (
	"os"
	"time"
)

func main() {
	duration, _ := time.ParseDuration(os.Args[1] + "s")
	time.Sleep(duration)
}
