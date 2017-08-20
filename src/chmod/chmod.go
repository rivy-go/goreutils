package main

import (
	"os"
	"strconv"
)

func main() {
	intMode, _ := strconv.Atoi(os.Args[1])
	mode := os.FileMode(intMode)
	os.Chmod(os.Args[2], mode)

}
