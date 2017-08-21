package main

import (
	"flag"
	"time"
)

func main() {
	flag.Parse()
	duration, _ := time.ParseDuration(flag.Arg(0) + "s")
	time.Sleep(duration)
}
