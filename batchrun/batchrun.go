package batchrun

import (
  "../datafetch"
  "../store"
  "net/http"
  "os"
  "errors"
  "fmt"
)

func PerformDump(dirName string) (error) {
  info, err := os.Stat(dirName)
  if err != nil {
    return err
  } else if !info.IsDir() {
    return errors.New(fmt.Sprintf("Expected directory, but was file: %+v", dirName))
  }

  client := &http.Client{}
  datums, err := datafetch.Fetch(client)
  if err != nil {
    return err
  }
  return store.StoreAll(dirName, datums)
}
