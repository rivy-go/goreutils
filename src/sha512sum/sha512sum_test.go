package main

import (
	"bytes"
	"crypto/sha512"
	"fmt"
	"io/ioutil"
	"os/exec"
	"testing"
)

func TestOutput(t *testing.T) {
	cmd := exec.Command("go", "run", "sha512sum.go", "../../Readme.md")
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Run()
	file, _ := ioutil.ReadFile("../../Readme.md")
	hash := sha512.Sum512(file)
	hashString := fmt.Sprintf("%x", hash)
	if out.String() != hashString+"\n" {
		t.Fail()
	}
}
