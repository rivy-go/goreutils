package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Parse()
	fmt.Print(flag.Parse(0))
}
