package main

import (
	"flag"
	"os"
)

func main() {
	flag.Parse()
	process, _ := os.FindProcess(os.Getppid())
	process.Kill()
}
