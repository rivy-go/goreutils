package main

import (
	"flag"
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
		proccess.Kill()
	} else {
		proccess.Signal(syscall.SIGTERM)
	}
}
