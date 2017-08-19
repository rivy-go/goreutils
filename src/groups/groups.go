package main

import (
	"fmt"
	"os"
	"os/user"
	"strings"
)

func main() {
	u, _ := user.Lookup(os.Args[1])
	groups, _ := u.GroupIds()
	var groupNames []string
	for _, element := range groups {
		group, _ := user.LookupGroupId(element)
		groupNames = append(groupNames, group.Name)
	}
	fmt.Println(strings.Join(groupNames, " "))
}
