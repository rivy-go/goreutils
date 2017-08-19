package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	var fileNames []string
	currentDir, _ := os.Getwd()
	files, _ := ioutil.ReadDir(currentDir)
	for _, file := range files {
		fileNames = append(fileNames, file.Name())
	}
	fmt.Println(strings.Join(fileNames, " "))
}
