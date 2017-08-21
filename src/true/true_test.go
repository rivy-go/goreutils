package main

import (
	"os/exec"
	"testing"
)

func TestOutput(t *testing.T) {
	cmd := exec.Command("go", "run", "true.go")
	err := cmd.Run()
	if err != nil {
		t.Fail()
	}
}
