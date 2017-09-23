package poloniextickerdumper

import (
  "testing"
  "google.golang.org/appengine/aetest"
  "google.golang.org/appengine/datastore"
)

func TestStoreTickerDatums(t *testing.T) {
  ctx, done, err := aetest.NewContext()
  if err != nil {
    t.Fatal(err)
  }
  defer done()

  original := TickerDatum{Timestamp: 12345, CurrencyPair: "ETH_BTC", Last: "123", LowestAsk: "234", HighestBid: "345", PercentChange: "456", BaseVolume: "567", QuoteVolume: "678", IsFrozen: "789", High24hr: "890", Low24hr: "901"}
  key, err := Store(ctx, &original)
  if err != nil {
    t.Fatal(err)
  }
  if key == nil {
    t.Fatal("Didn't receive a key")
  }
  var stored TickerDatum
  if err := datastore.Get(ctx, key, &stored); err != nil {
    t.Fatal(err)
  }
  if stored != original {
    t.Fatalf("Expected all struct fields to match, instead got %+v", stored)
  }
}
