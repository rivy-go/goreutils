package main

import (
	"bytes"
	"os/exec"
	"runtime"
	"testing"
)

func TestOutput(t *testing.T) {
	cmd := exec.Command("go", "run", "arch.go")
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Run()
	if out.String() != runtime.GOARCH+"\n" {
		t.Fail()
	}
}
