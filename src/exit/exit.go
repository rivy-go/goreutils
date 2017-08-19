package main

import "os"

func main() {
	process, _ := os.FindProcess(os.Getppid())
	process.Kill()
}
