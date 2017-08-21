package main

import (
	"bytes"
	"os/exec"
	"testing"
  "os/user"
  "strings"
)

func TestOutput(t *testing.T) {
  testUser, _ := user.Current()
	cmd := exec.Command("go", "run", "groups.go", testUser.Username)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Run()
  u, _ := user.Lookup(testUser.Username)
	groups, _ := u.GroupIds()
	var groupNames []string
	for _, element := range groups {
		group, _ := user.LookupGroupId(element)
		groupNames = append(groupNames, group.Name)
	}
	output := strings.Join(groupNames, " ")
	if out.String() != output {
		t.Fail()
	}
}
