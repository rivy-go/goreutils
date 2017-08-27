package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	flag.Parse()
	var fileNames []string
	currentDir, _ := os.Getwd()
	files, err := ioutil.ReadDir(currentDir)
	for _, file := range files {
		fileNames = append(fileNames, file.Name())
	}
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	fmt.Print(strings.Join(fileNames, " "))
}
