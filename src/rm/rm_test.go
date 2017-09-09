package main

import (
	"errors"
	"os"
	"os/exec"
	"testing"
)

func TestDeleted(t *testing.T) {
	file, err := os.Create("test")
	if err != nil {
		t.Fatal(errors.New("have no access to filesystem"))
	}
	file.Close()
	cmd := exec.Command("go", "run", "rm.go", "test")
	cmd.Run()
	_, err = os.Stat("test")
	if err == nil {
		t.Fail()
	}
}
