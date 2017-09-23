package main

import (
	"flag"
	"fmt"
	"runtime"
)

func main() {
	flag.Parse()
	fmt.Println(runtime.GOARCH)
}
