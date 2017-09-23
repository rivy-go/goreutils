package main

import (
	"os"
	"os/exec"
	"testing"
)

func TestCreated(t *testing.T) {
	cmd := exec.Command("go", "run", "mkdir.go", "test")
	cmd.Run()
	stat, err := os.Stat("test")
	if err != nil {
		t.Fail()
		return
	}
	if !stat.IsDir() {
		t.Fail()
	}
	os.Remove("test")
}
