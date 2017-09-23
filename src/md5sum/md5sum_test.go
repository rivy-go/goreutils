package main

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"os/exec"
	"testing"
)

func TestOutput(t *testing.T) {
	cmd := exec.Command("go", "run", "md5sum.go", "../../Readme.md")
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Run()
	file, _ := ioutil.ReadFile("../../Readme.md")
	hash := md5.Sum(file)
	hashString := fmt.Sprintf("%x", hash)
	if out.String() != hashString+"\n" {
		t.Fail()
	}
}
