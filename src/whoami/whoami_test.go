package main

import (
	"bytes"
	"os/exec"
	"os/user"
	"testing"
)

func TestOutput(t *testing.T) {
	cmd := exec.Command("go", "run", "whoami.go")
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Run()
	currentUser, _ := user.Current()
	if out.String() != currentUser.Username {
		t.Fail()
	}
}
