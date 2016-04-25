package main

import (
    "fmt"
    "time"
)

func emit(chanChannel chan chan string, done chan bool) {
  wordChannel := make(chan string)
  chanChannel <- wordChannel
  words := []string{"The", "quick", "brown", "fox"}
  i := 0

  defer close(wordChannel)

  t := time.NewTimer(3 * time.Second)

  for {
    select {
    case wordChannel <- words[i]:
      i += 1
      if i == len(words) {
        i = 0
      }
    case <- done:
      fmt.Println("Received done message")
      close(done)
      return
    case <- t.C:
      return
    }
  }
}

func main() {
  channelCh := make(chan chan string)
  doneChannel := make(chan bool)
  
  // concurrently run emit
  go emit(channelCh, doneChannel)

  wordChannel := <- channelCh

  for word := range wordChannel {
    fmt.Printf("%s ", word)
  }

}
