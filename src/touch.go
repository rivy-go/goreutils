package main

import (
  "os"
  "time"
)

func main() {
  current_time := time.Now()
  os.Chtimes(os.Args[1], current_time, current_time)
}
