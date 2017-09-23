package main

import (
	"bytes"
	"os/exec"
	"testing"
)

func TestOutput(t *testing.T) {
	cmd := exec.Command("go", "run", "echo.go", "testing")
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Run()
	if out.String() != "testing"+"\n" {
		t.Fail()
	}
}
