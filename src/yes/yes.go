package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Parse()
	var arg = "yes"
	if flag.Arg(0) != "" {
		arg = flag.Arg(0)
	}
	for true {
		fmt.Println(arg)
	}
}
