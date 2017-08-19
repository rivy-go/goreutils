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
	var group_names []string
	for _, element := range groups {
		group, _ := user.LookupGroupId(element)
		group_names = append(group_names, group.Name)
	}
	fmt.Println(strings.Join(group_names, " "))
}
