package main

import (
  "log"
  "os"
  "encoding/csv"
)

func ReadCsv(filename string) {
  f, err := os.Open(filename)
  if err != nil {
    log.Fatal(err)
  }

  defer f.Close()
  r := csv.NewReader(f)
  r.Read() // Do not include the headers

  // Read All the rows into memory at once
  recs, err := r.ReadAll()
  if err != nil {
    log.Fatal(err)
  }

  for _, row := range recs {
    printRow(row)
  }

}

func printRow(row []string) {
  log.Printf("len(row) %d\n", len(row))
  for i, col := range row {
    log.Printf("[%d]: %s\n", i, col)
  }
}
