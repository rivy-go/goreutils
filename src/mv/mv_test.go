package main

import (
	"errors"
	"io/ioutil"
	"os"
	"os/exec"
	"testing"
)

func TestMoved(t *testing.T) {
	file, err := os.Create("old")
	if err != nil {
		t.Fatal(errors.New("have no access to filesystem"))
	}
	file.WriteString("test")
	file.Close()
	cmd := exec.Command("go", "run", "mv.go", "old", "new")
	cmd.Run()
	newFile, err := ioutil.ReadFile("new")
	if err != nil {
		t.Fail()
	}
	if string(newFile) != "test" {
		t.Fail()
	}
	_, err = os.Stat("old")
	if err == nil {
		t.Fail()
	}
	os.Remove("new")
}
