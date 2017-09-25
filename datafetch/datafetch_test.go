package datafetch

import (
  "testing"
  "net/http"
  "io"
  "bytes"
)

func TestParsing(t *testing.T) {
  str := []byte("{\"BTC_BCN\":{\"id\":7,\"last\":\"0.00000031\",\"lowestAsk\":\"0.00000031\",\"highestBid\":\"0.00000030\",\"percentChange\":\"-0.03125000\",\"baseVolume\":\"40.22111561\",\"quoteVolume\":\"129435889.20216185\",\"isFrozen\":\"0\",\"high24hr\":\"0.00000032\",\"low24hr\":\"0.00000030\"},\"BTC_BELA\":{\"id\":8,\"last\":\"0.00003598\",\"lowestAsk\":\"0.00003599\",\"highestBid\":\"0.00003598\",\"percentChange\":\"0.16628849\",\"baseVolume\":\"29.74189120\",\"quoteVolume\":\"908673.62756949\",\"isFrozen\":\"0\",\"high24hr\":\"0.00003747\",\"low24hr\":\"0.00003000\"},\"BTC_BLK\":{\"id\":10,\"last\":\"0.00004560\",\"lowestAsk\":\"0.00004545\",\"highestBid\":\"0.00004519\",\"percentChange\":\"-0.07990314\",\"baseVolume\":\"40.01428327\",\"quoteVolume\":\"836047.58947475\",\"isFrozen\":\"0\",\"high24hr\":\"0.00005217\",\"low24hr\":\"0.00004327\"}}")
  tickers, err := ParseResponse(str)
  if err != nil {
    t.Fatalf("Expected to have no errors but got: %+v", err)
  }
  if len(tickers) != 3 {
    t.Fatalf("Expected to have three tickers but had: %+v", len(tickers))
  }
  var BTC_BCN *TickerDatum
  for i := range tickers {
    var ticker = tickers[i]
    if ticker.CurrencyPair == "BTC_BCN" {
      BTC_BCN = ticker
    }
  }
  if BTC_BCN == nil {
    t.Fatal("Expected to find BTC_BCN ticker, but didn't")
  }
  if BTC_BCN.Last != "0.00000031" {
    t.Fatal("Expected to have correct value for BTC_BCN.Last, but got:", BTC_BCN.Last)
  }
}

func TestFetchDataWithErrorStatusCode(t *testing.T) {
  _, err := FetchData(GetterMock{ get: func (url string) (*http.Response, error) {
    return &http.Response { Body: ResponseBody(""), StatusCode: 400 },nil
  }})
  if err == nil {
    t.Fatal("Expected to get error")
  }
}

func TestFetchDataWithBody(t *testing.T) {
  bodyBytes, err := FetchData(GetterMock{ get: func (url string) (*http.Response, error) {
    return &http.Response { Body: ResponseBody("foo"), StatusCode: 200 },nil
  }})
  if err != nil {
    t.Fatal("Shouldn't have errors")
  } else if !bytes.Equal(bodyBytes, []byte("foo")) {
    t.Fatal("Body not passed correctly")
  }
}

type GetterMock struct {
  get func(string) (*http.Response, error)
}

func (g GetterMock) Get(url string) (*http.Response, error) { return g.get(url) }

type nopCloser struct {
    io.Reader
}

func (nopCloser) Close() error { return nil }

func ResponseBody(str string) io.ReadCloser {
  return nopCloser{bytes.NewBufferString(str)}
}
