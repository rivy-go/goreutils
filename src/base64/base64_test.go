package main

import (
	"bytes"
	"encoding/base64"
	"io/ioutil"
	"os/exec"
	"testing"
)

func TestOutput(t *testing.T) {
	cmd := exec.Command("go", "run", "base64.go", "../../Readme.md")
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Run()
	file, _ := ioutil.ReadFile("../../Readme.md")
	hash := base64.StdEncoding.EncodeToString(file)
	if out.String() != hash+"\n" {
		t.Fail()
	}
}
