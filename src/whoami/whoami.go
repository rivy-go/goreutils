package main

import (
	"flag"
	"fmt"
	"log"
	"os/user"
)

func main() {
	flag.Parse()
	currentUser, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(currentUser.Username)
}
