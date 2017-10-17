package main

import (
	"bytes"
	"encoding/base32"
	"io/ioutil"
	"os/exec"
	"testing"
)

func TestOutput(t *testing.T) {
	cmd := exec.Command("go", "run", "base32.go", "../../Readme.md")
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Run()
	file, _ := ioutil.ReadFile("../../Readme.md")
	hash := base32.StdEncoding.EncodeToString(file)
	if out.String() != hash+"\n" {
		t.Fail()
	}
}
