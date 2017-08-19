package main

import (
	"fmt"
	"os/user"
)

func main() {
	current_user, _ := user.Current()
	fmt.Println(current_user.Username)
}
