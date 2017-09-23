package poloniextickerdumper

import (
  "testing"
  "google.golang.org/appengine/aetest"
  "google.golang.org/appengine/datastore"
)

func TestFullBatch(t *testing.T) {
  ctx, done, err := aetest.NewContext()
  if err != nil {
    t.Fatal(err)
  }
  defer done()

  count, err := datastore.NewQuery(DATASTORE_ENTITY_TYPE).Count(ctx)

  if err != nil {
    t.Fatal(err)
  } else if count > 0 {
    t.Fatal("Expected datastore entity count to be zero, instead got:", count)
  }

  keys, err := PerformDump(ctx)
  if err != nil {
    t.Fatal(err)
  } else if len(keys) < 2 {
    t.Fatal("Expected to store at least two keys, instead stored:", len(keys))
  }

}
