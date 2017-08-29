package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Parse()
	for _, env := range os.Environ() {
		fmt.Println(env)
	}
}
