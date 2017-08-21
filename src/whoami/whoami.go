package main

import (
	"flag"
	"fmt"
	"os/user"
)

func main() {
	flag.Parse()
	currentUser, _ := user.Current()
	fmt.Print(currentUser.Username)
}
