package main

import (
	"bytes"
	"io/ioutil"
	"os/exec"
	"testing"
)

func TestOutput(t *testing.T) {
	cmd := exec.Command("go", "run", "cat.go", "../../Readme.md")
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Run()
	file, _ := ioutil.ReadFile("../../Readme.md")
	if out.String() != string(file) {
		t.Fail()
	}
}
