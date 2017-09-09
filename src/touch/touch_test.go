package main

import (
	"os"
	"os/exec"
	"testing"
)

func TestCreated(t *testing.T) {
	cmd := exec.Command("go", "run", "touch.go", "test")
	cmd.Run()
	_, err := os.Stat("test")
	if err != nil {
		t.Fail()
	}
	os.Remove("test")
}
