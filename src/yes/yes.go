package main

import (
	"flag"
	"fmt"
)

func main() {
	for true {
		flag.Parse()
		fmt.Println(flag.Arg(0))
	}
}
