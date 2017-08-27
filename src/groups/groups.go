package main

import (
	"flag"
	"fmt"
	"os"
	"os/user"
	"strings"
)

func main() {
	flag.Parse()
	u, _ := user.Lookup(flag.Arg(0))
	groups, _ := u.GroupIds()
	var groupNames []string
	for _, element := range groups {
		group, err := user.LookupGroupId(element)
		if err != nil {
			fmt.Print(err)
			os.Exit(1)
		}
		groupNames = append(groupNames, group.Name)
	}
	fmt.Print(strings.Join(groupNames, " "))
}
