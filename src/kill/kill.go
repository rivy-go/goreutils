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
	pid, _ := strconv.Atoi(flag.Arg(0))
	proccess, _ := os.FindProcess(pid)
	if killProc {
		err := proccess.Kill()
		if err != nil {
			fmt.Print("Error")
			os.Exit(1)
		}
	} else {
		err := proccess.Signal(syscall.SIGTERM)
		if err != nil {
			fmt.Print("Error")
			os.Exit(1)
		}
	}
}
