package main

import (
  "fmt"
)

type  SummableSlice [] int

func (s SummableSlice) sum() int {
  sum := 0
  for _, i := range s {
    sum += i
  }
  return sum
}

func main() {
  var s SummableSlice = SummableSlice{1,1,2,3,5,8,13}
  fmt.Printf("%d\n", s.sum())
}

