package main

import (
	"flag"
	"fmt"
	"os"
	"os/user"
)

func main() {
	flag.Parse()
	currentUser, err := user.Current()
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	fmt.Print(currentUser.Username)
}
