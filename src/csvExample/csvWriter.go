package main

import (
  "log"
  "os"
  "encoding/csv"
)

func WriteCsv(filename string) {
  f, err := os.Create(filename)
  if err != nil {
    log.Fatal(err)
  }

  defer f.Close()
  w := csv.NewWriter(f)
  w.Write([]string{"first", "last", "email"})
  w.Flush()
  w.Write([]string{"John", "Doe", "john.doe@example,com"})
  w.Flush()
}
