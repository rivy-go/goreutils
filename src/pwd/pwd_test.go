package main

import (
	"bytes"
	"os"
	"os/exec"
	"testing"
)

func TestOutput(t *testing.T) {
	cmd := exec.Command("go", "run", "pwd.go")
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Run()
	workdir, _ := os.Getwd()
	if out.String() != workdir {
		t.Fail()
	}
}
