package main

import (
	"flag"
	"os"
	"strconv"
)

func main() {
	flag.Parse()
	pid, _ := strconv.Atoi(flag.Arg(0))
	proccess, _ := os.FindProcess(pid)
	proccess.Kill()
}
