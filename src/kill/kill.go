package main

import (
	"flag"
	"fmt"
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
		fmt.Print(err)
		os.Exit(1)
	}
	proccess, err := os.FindProcess(pid)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	if killProc {
		err := proccess.Kill()
		if err != nil {
			fmt.Print(err)
			os.Exit(1)
		}
	} else {
		err := proccess.Signal(syscall.SIGTERM)
		if err != nil {
			fmt.Print(err)
			os.Exit(1)
		}
	}
}
