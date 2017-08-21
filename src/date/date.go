package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	flag.Parse()
	fmt.Print(time.Now())
}
