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
	fmt.Print(currentUser.Username)
	if err != nil {
		fmt.Print("Error")
		os.Exit(1)
	}
}
