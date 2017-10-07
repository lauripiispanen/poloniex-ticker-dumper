package main

import (
  "os"
  "flag"
  "fmt"
  "time"
  "log"
  "net/http"
  "github.com/lauripiispanen/poloniex-ticker-dumper/store"
  "github.com/lauripiispanen/poloniex-ticker-dumper/datafetch"
)

func main() {
  dirName := getDirnameFlagOrExit()
  datumChan := make(chan []*datafetch.TickerDatum)
  go func() {
    for datums := range datumChan {
      err := store.StoreAll(dirName, datums)
      if err != nil {
        log.Println("Error storing datums:", err)
      }
    }
  }()

  ticker := time.NewTicker(2 * time.Second)
  for _ = range ticker.C {
    go func() {
      log.Println("Fetching...")
      client := &http.Client{}
      datums, err := datafetch.Fetch(client)
      if err != nil {
        log.Println("Error fetching:", err)
      } else {
        datumChan <- datums
      }
    }()
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
