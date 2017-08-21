package main

import (
	"flag"
	"fmt"
	"runtime"
)

func main() {
	flag.Parse()
	fmt.Print(runtime.GOARCH)
}
