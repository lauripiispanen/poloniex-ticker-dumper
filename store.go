package poloniextickerdumper

import (
  "google.golang.org/appengine/datastore"
  "golang.org/x/net/context"
  "errors"
)


func Store(ctx context.Context, datum *TickerDatum) (*datastore.Key, error) {
  k, err := StoreAll(ctx, []*TickerDatum{datum})
  if err != nil {
    return nil, err
  } else if len(k) == 1 {
    return k[0], nil
  } else {
    return nil, errors.New("Got multiple keys from single store")
  }
}

func StoreAll(ctx context.Context, datums []*TickerDatum) ([]*datastore.Key, error) {
  var keys []*datastore.Key
  for range datums {
    keys = append(keys, datastore.NewIncompleteKey(ctx, DATASTORE_ENTITY_TYPE, nil))
  }
  k, err := datastore.PutMulti(ctx, keys, datums)
  if err != nil {
    return nil, err
  } else {
    return k, nil
  }
}

const DATASTORE_ENTITY_TYPE = "TickerDatum"
