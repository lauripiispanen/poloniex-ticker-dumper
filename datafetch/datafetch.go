package datafetch

import (
  "net/http"
  "encoding/json"
  "time"
  "io/ioutil"
  "errors"
)

func Fetch(client *http.Client) ([]*TickerDatum, error) {
  bodyBytes, err := FetchData(client)
  if err != nil {
    return nil, err
  }
  return ParseResponse(bodyBytes)
}

func FetchData(getter Getter) ([]byte, error) {
  resp, err := getter.Get("https://poloniex.com/public?command=returnTicker")
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()
  if resp.StatusCode == 200 {
    bodyBytes, err := ioutil.ReadAll(resp.Body)
    return bodyBytes, err
  } else {
    return nil, errors.New("Got non-200 status")
  }
}

func ParseResponse(input []byte) ([]*TickerDatum, error) {
  var parsed map[string]TickerData
  var out []*TickerDatum
  err := json.Unmarshal(input, &parsed)
  if err != nil {
    return nil, err
  }
  timestamp := time.Now().Unix()

  for key, value := range parsed {
    datum := TickerDatum {Timestamp: timestamp, CurrencyPair: key, Last: value.Last, LowestAsk: value.LowestAsk, HighestBid: value.HighestBid, PercentChange: value.PercentChange, BaseVolume: value.BaseVolume, QuoteVolume: value.QuoteVolume, IsFrozen: value.IsFrozen, High24hr: value.High24hr, Low24hr: value.Low24hr }
    out = append(out, &datum)
  }
  return out, nil
}


type Getter interface {
  Get(url string) (resp *http.Response, err error)
}

type TickerData struct {
  Id            int
  Last          string
  LowestAsk     string
  HighestBid    string
  PercentChange string
  BaseVolume    string
  QuoteVolume   string
  IsFrozen      string
  High24hr      string
  Low24hr       string
}

type TickerDatum struct {
  Timestamp     int64
  CurrencyPair  string
  Last          string
  LowestAsk     string
  HighestBid    string
  PercentChange string
  BaseVolume    string
  QuoteVolume   string
  IsFrozen      string
  High24hr      string
  Low24hr       string
}
