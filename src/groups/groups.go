package main

import (
	"flag"
	"fmt"
	"log"
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
			log.Fatal(err)
		}
		groupNames = append(groupNames, group.Name)
	}
	fmt.Print(strings.Join(groupNames, "\n"))
}
