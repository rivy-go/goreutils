package main

import (
	"fmt"
	"os"
)

func main() {
	workdir, _ := os.Getwd()
	fmt.Println(workdir)
}