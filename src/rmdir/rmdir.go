package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	flag.Parse()
	stat, err := os.Stat(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}
	if !stat.IsDir() {
		log.Fatal("is no dir")
	}
	files, err := ioutil.ReadDir(flag.Arg(0))
	if len(files) != 0 {
		log.Fatal("not empty")
	}
	err = os.Remove(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}
}
