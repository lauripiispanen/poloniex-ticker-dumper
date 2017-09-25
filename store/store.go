package store

import (
  "../datafetch"
  "path"
  "time"
  "fmt"
  "bufio"
  "os"
  "io"
  "strconv"
  "encoding/csv"
)


func Store(dirName string, datum *datafetch.TickerDatum) error {
  return StoreAll(dirName, []*datafetch.TickerDatum{datum})
}

func StoreAll(dirName string, datums []*datafetch.TickerDatum) error {
  for i := range datums {
    datum := datums[i]
    f, isNewFile, err := FileForDatum(dirName, datum)
    if err != nil {
      return err
    }
    if err = WriteDatum(f, isNewFile, datum); err != nil {
		  return err
	  }
  }
  return nil
}

func FileForDatum(dirName string, datum *datafetch.TickerDatum) (*os.File, bool, error) {
  fileName := StorableFilePath(dirName, datum)
  _, err := os.Stat(fileName)
  isNewFile := os.IsNotExist(err)

  f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
  if err != nil {
    return nil, isNewFile, err
  } else {
    return f, isNewFile, nil
  }
}

func WriteDatum(writer io.Writer, writeHeader bool, datum *datafetch.TickerDatum) error {
  CSV_HEADERS := []string{ "Timestamp", "CurrencyPair", "Last", "LowestAsk", "HighestBid", "PercentChange", "BaseVolume", "QuoteVolume", "IsFrozen", "High24hr", "Low24hr" }
  fw := bufio.NewWriter(writer)
  w := csv.NewWriter(fw)
  if writeHeader {
    w.Write(CSV_HEADERS)
  }
  w.Write(ToRecord(datum))
  w.Flush()
  if err := w.Error(); err != nil {
    return err
  }
  return nil
}

func StorableFilePath(dirName string, datum *datafetch.TickerDatum) string {
  t := time.Unix(datum.Timestamp, 0).UTC()
  timeStr := fmt.Sprintf("%d-%d-%d.csv", int(t.Year()), int(t.Month()), int(t.Day()))
  return path.Join(dirName, timeStr)
}

func ToRecord(datum *datafetch.TickerDatum) []string {
  return []string{ strconv.FormatInt(datum.Timestamp, 10), datum.CurrencyPair, datum.Last, datum.LowestAsk, datum.HighestBid, datum.PercentChange, datum.BaseVolume, datum.QuoteVolume, datum.IsFrozen, datum.High24hr, datum.Low24hr }
}

const DATASTORE_ENTITY_TYPE = "TickerDatum"
