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
	fmt.Print(strings.Join(fileNames, " "))
	if err != nil {
		fmt.Print("Error")
		os.Exit(1)
	}
}
