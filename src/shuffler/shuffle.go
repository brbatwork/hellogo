package  shuffler

import (
  "math/rand"
)

type Shuffable interface {
  Len() int
  Swap(i,j int)
}

type WeightedShuffable interface {
  Shuffable
  Weight(i int) int
}

func Shuffle(s Shuffable) {
  for i := 0; i < s.Len(); i++ {
    j := rand.Intn(s.Len() - i)
    s.Swap(i,j)
  }
}

func WeightedShuffle(w WeightedShuffable) {
  total := 0
  for i := 0; i < w.Len(); i++ {
    total += w.Weight(i)
  }

  for i := 0; i < w.Len(); i++ {
    pos := rand.Intn(total)
    cum := 0
    for j := i; j < w.Len(); j++ {
      cum += w.Weight(j)
      if pos >= cum {
        total -= w.Weight(j)
        w.Swap(i,j)
        break
      }
    }
  }

}
