package main

import (
    "fmt"
    "time"
    "math/rand"
)

func reader(ch chan int) {
  t := time.NewTimer(10 * time.Second)

  for {
    select {
    case i := <-ch:
      fmt.Printf("%d\n", i)
    case <- t.C:
      ch = nil
    }
  }
}

func writer(ch chan int) {
  stopper := time.NewTimer(2 * time.Second)
  restarter := time.NewTimer(5 * time.Second)

  savedCh := ch

  for {
    select  {
    case ch <- rand.Intn(100):
    case <- stopper.C:
      ch = nil
    case <- restarter.C:
      ch = savedCh
    }
  }
}


func main() {
  ch := make(chan int)
  go reader(ch)
  go writer(ch)
  time.Sleep(30 * time.Second)


}
