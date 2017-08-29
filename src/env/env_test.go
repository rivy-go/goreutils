package main

import (
	"bytes"
	"os"
	"os/exec"
	"testing"
)

func TestOutput(t *testing.T) {
	cmd := exec.Command("go", "run", "env.go")
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Run()
	var envs string
	for _, env := range os.Environ() {
		envs = envs + env + "\n"
	}
	if out.String() != envs {
		t.Fail()
	}
}
