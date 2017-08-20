package main

import (
	"os"
)

func main() {
	os.Rename(os.Args[1], os.Args[2])
}
