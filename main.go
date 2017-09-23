package poloniextickerdumper

import (
  "fmt"
  "net/http"
  "google.golang.org/appengine"
)

func init() {
  http.HandleFunc("/run", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
  keys, err := PerformDump(appengine.NewContext(r))
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  fmt.Fprint(w, "Count:", len(keys))
}
