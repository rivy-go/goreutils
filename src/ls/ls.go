package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
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
		log.Fatal(err)
	}
	fmt.Println(strings.Join(fileNames, "\n"))
}
