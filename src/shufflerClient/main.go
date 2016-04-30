package main

import (
  "fmt"
  "shuffler"
)

type weightString struct {
  weight int
  as string
}

// Shuffable ints
type intSlice []int
func (is intSlice) Len() int {
  return len(is)
}

func (is intSlice) Swap(i, j int) {
  is[i], is[j] = is[j], is[i]
}

// Shuffable strings
type stringSlice []weightString
func (ss stringSlice) Len() int {
  return len(ss)
}

func (ss stringSlice) Swap(i, j int) {
  ss[i], ss[j] = ss[j], ss[i]
}

func (ss stringSlice) Weight(i int) int {
  return ss[i].weight
}



func main() {
  is := intSlice{1,2,3,4,5,6}
  shuffler.Shuffle(is)
  fmt.Printf("%v\n", is)

  ss := stringSlice{weightString{100, "hello"}, weightString{200, "world"}, weightString{10, "goodbye"}}
  shuffler.Shuffle(ss)
  fmt.Printf("%v\n", ss)
  shuffler.WeightedShuffle(ss)
  fmt.Printf("%v\n", ss)
}
