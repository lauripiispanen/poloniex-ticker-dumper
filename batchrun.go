package poloniextickerdumper

import (
  "golang.org/x/net/context"
  "google.golang.org/appengine/datastore"
  "google.golang.org/appengine/urlfetch"
)

func PerformDump(ctx context.Context) ([]*datastore.Key, error) {
  client := urlfetch.Client(ctx)
  datums, err := Fetch(client)
  if err != nil {
    return nil, err
  }
  var storable []*TickerDatum
  for i := range datums {
    datum := datums[i]
    if datum.CurrencyPair == "BTC_ETH" {
      storable = append(storable, &datums[i])
    }
  }
  return StoreAll(ctx, storable)
}
