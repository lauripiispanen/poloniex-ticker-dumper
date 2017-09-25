package main

import (
  "os"
  "os/signal"
  "syscall"
  "sync"
  "flag"
  "fmt"
  "./batchrun"
)

func main() {
  dirName := getDirnameFlagOrExit()
  callback := func() {
    batchrun.PerformDump(dirName)
  }

  var wg sync.WaitGroup
  wg.Add(1)
  go listenForSIGUSR1(wg, callback)
  wg.Wait()
}

func listenForSIGUSR1(wg sync.WaitGroup, callback func()) {
  defer wg.Done()
  c := make(chan os.Signal, 1)
  signal.Notify(c, syscall.SIGUSR1)
  for _ = range c {
    callback()
  }
}

func getDirnameFlagOrExit() string {
  dirNamePtr := flag.String("dir", "", "directory to use")
  flag.Parse()
  if *dirNamePtr == "" {
    fmt.Println("missing flag 'dir'")
    os.Exit(1)
  }
  return *dirNamePtr
}
