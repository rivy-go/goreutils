package main

import (
	"flag"
	"fmt"
	"os/user"
	"strings"
)

func main() {
	flag.Parse()
	u, _ := user.Lookup(flag.Arg(0))
	groups, _ := u.GroupIds()
	var groupNames []string
	for _, element := range groups {
		group, _ := user.LookupGroupId(element)
		groupNames = append(groupNames, group.Name)
	}
	fmt.Print(strings.Join(groupNames, " "))
}
