package main

import (
	"bytes"
	"io/ioutil"
	"os/exec"
	"testing"
)

func TestCreated(t *testing.T) {
	file, _ := ioutil.ReadFile("test")
	cmd := exec.Command("go", "touch", "test")
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Run()
	if out.String() != string(file) {
		t.Fail()
	}
}
