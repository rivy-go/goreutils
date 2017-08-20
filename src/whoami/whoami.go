package main

import (
	"fmt"
	"os/user"
)

func main() {
	currentUser, _ := user.Current()
	fmt.Print(currentUser.Username)
}
