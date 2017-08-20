package main

import (
	"os"
	"strconv"
)

func main() {
	pid, _ := strconv.Atoi(os.Args[1])
	proccess, _ := os.FindProcess(pid)
	proccess.Kill()
}
