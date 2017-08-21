package main

import (
	"bytes"
	"os/exec"
	"testing"
  "os"
  "io/ioutil"
  "strings"
)

func TestOutput(t *testing.T) {
	cmd := exec.Command("go", "run", "ls.go")
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Run()
  var fileNames []string
	currentDir, _ := os.Getwd()
	files, _ := ioutil.ReadDir(currentDir)
	for _, file := range files {
		fileNames = append(fileNames, file.Name())
	}
	output := strings.Join(fileNames, " ")
	if out.String() != output {
		t.Fail()
	}
}
