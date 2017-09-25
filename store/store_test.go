package store

import (
  "../datafetch"
  "testing"
  "bytes"
)

func TestStorableFilePath(t *testing.T) {
  path := StorableFilePath("foo", CreateTestDatum())
  expected := "foo/2009-2-13.csv"
  if path != expected {
    t.Fatalf("Expected path to equal %+v, instead got %+v", expected, path)
  }
}

func CreateTestDatum() *datafetch.TickerDatum {
  return &datafetch.TickerDatum{Timestamp: 1234567890, CurrencyPair: "ETH_BTC", Last: "123", LowestAsk: "234", HighestBid: "345", PercentChange: "456", BaseVolume: "567", QuoteVolume: "678", IsFrozen: "789", High24hr: "890", Low24hr: "901"}
}

func TestWriteDatumWithHeaders(t *testing.T) {
  buf := new(bytes.Buffer)
  WriteDatum(buf, true, CreateTestDatum())
  result := buf.String()
  expected := "Timestamp,CurrencyPair,Last,LowestAsk,HighestBid,PercentChange,BaseVolume,QuoteVolume,IsFrozen,High24hr,Low24hr\n1234567890,ETH_BTC,123,234,345,456,567,678,789,890,901\n"
  if result != expected {
    t.Fatalf("Expected result to equal %+v, instead got %+v", expected, result)
  }
}

func TestWriteDatumWithoutHeaders(t *testing.T) {
  buf := new(bytes.Buffer)
  WriteDatum(buf, false, CreateTestDatum())
  result := buf.String()
  expected := "1234567890,ETH_BTC,123,234,345,456,567,678,789,890,901\n"
  if result != expected {
    t.Fatalf("Expected result to equal %+v, instead got %+v", expected, result)
  }
}
