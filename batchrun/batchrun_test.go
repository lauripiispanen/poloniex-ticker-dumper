package batchrun

import (
  "testing"
  "os"
  "io/ioutil"
)

func TestFullBatch(t *testing.T) {
  dir, err := ioutil.TempDir("", "example")
  if err != nil {
    t.Fatalf("Error creating temp dir", err)
  }
  defer os.RemoveAll(dir)

  dirFile, err := os.Open(dir)
  if err != nil {
    t.Fatalf("Error opening temp dir", err)
  }
  originalFileNames, err := dirFile.Readdirnames(0)
  if err != nil {
    t.Fatalf("Error reading temp dir", err)
  }

  PerformDump(dir)

  fileNames, err := dirFile.Readdirnames(0)
  if err != nil {
    t.Fatalf("Error reading temp dir", err)
  }

  if len(fileNames) < len(originalFileNames) {
    t.Fatalf("Expected number of files in output directory to increase")
  }
}
