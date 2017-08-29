package main

import (
	"flag"
	"log"
	"os"
	"strconv"
	"syscall"
)

func main() {
	var killProc bool
	flag.BoolVar(&killProc, "kill", false, "Send a SIGKILL instead of a SIGTERM.")
	flag.Parse()
	pid, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}
	proccess, err := os.FindProcess(pid)
	if err != nil {
		log.Fatal(err)
	}
	if killProc {
		err := proccess.Kill()
		if err != nil {
			log.Fatal(err)
		}
	} else {
		err := proccess.Signal(syscall.SIGTERM)
		if err != nil {
			log.Fatal(err)
		}
	}
}
