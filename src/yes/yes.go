package main

import (
	"flag"
	"fmt"
)

func main() {
	for true {
		flag.Parse()
		fmt.Print(flag.Arg(0))
	}
}
