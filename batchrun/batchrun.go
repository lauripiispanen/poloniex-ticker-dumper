package batchrun

import (
  "golang.org/x/net/context"
  "google.golang.org/appengine/datastore"
  "google.golang.org/appengine/urlfetch"
  "../datafetch"
  "../store"
)

func Perform(ctx context.Context) ([]*datastore.Key, error) {
  client := urlfetch.Client(ctx)
  datums, err := datafetch.Fetch(client)
  if err != nil {
    return nil, err
  }
  var storable []*datafetch.TickerDatum
  for i := range datums {
    storable = append(storable, &datums[i])
  }
  return store.StoreAll(ctx, storable)
}
